package main

import (
	"api/internal/store"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"go.mongodb.org/mongo-driver/mongo"
)

type MidtransReponse struct {
	Token string `json:"token,omitempty"`
	URL string 	`json:"redirect_url,omitempty"`
} 

func (app *application) CartPaymentHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	order_id, userCart, err := app.store.Transaction.GetUserCart(r.Context(), user.ID)
	if err == mongo.ErrNoDocuments {
		app.badRequestResponse(w, r, errors.New("there are no products in user cart"))
		return
	}
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}	

	var total int
	for _, product := range userCart {
		total += int(product.Price) * int(product.Quantity)
	}

	// Request to midtrans
	rBody := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: order_id,
			GrossAmt: int64(total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	jsonBody, err := json.Marshal(rBody)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	
	req, err := http.NewRequest("POST", app.config.payment.paymentURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	authString :=  base64.StdEncoding.EncodeToString([]byte(app.config.payment.serverKey))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+authString)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	
	if err := app.store.Transaction.MarkTransaction(r.Context(), order_id); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	var res MidtransReponse
	if err := json.Unmarshal(body, &res); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, res)

}



func (app *application) ProductPaymentHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	var payload store.Cartproduct
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	order := store.Order {
		UserID: user.ID,
		Date: time.Now(),
		Status: "completed",
		Products: []store.Cartproduct{payload},
		Address: user.Addrres,
	}

	transaction_id, err := app.store.Transaction.CreateUserOrder(r.Context(), order)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	rBody := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: transaction_id,
			GrossAmt: int64(payload.Price) * int64(payload.Quantity),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	jsonBody, err := json.Marshal(rBody)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	
	req, err := http.NewRequest("POST", app.config.payment.paymentURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	authString :=  base64.StdEncoding.EncodeToString([]byte(app.config.payment.serverKey))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+authString)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	var res MidtransReponse
	if err := json.Unmarshal(body, &res); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, res)
}
