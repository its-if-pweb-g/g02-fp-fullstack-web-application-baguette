package main

import (
	"api/internal/store"
	"errors"
	"net/http"
)


func (app *application) UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) UserAddressHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user.Addrres); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) UserCartHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	cart, err := app.store.Transaction.GetUserCart(r.Context(), user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cart); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) AddProductToCartHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := app.store.Transaction.AddProduct(r.Context(), &payload, user.ID); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// func (app *application) PaymentHandler(w http.ResponseWriter, r *http.Request) {

// }

