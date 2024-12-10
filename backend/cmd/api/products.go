package main

import (
	"api/internal/store"
	"encoding/base64"
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

type ProductPayload struct {
	Name              string    `json:"name"`
	HeaderDescription string    `json:"header_description"`
	Description       string    `json:"description"`
	Price             int       `json:"price"`
	Stock             int       `json:"stock"`
	Sold              int       `json:"sold"`
	Image             string    `json:"image"`
	CreatedAt         time.Time `json:"created_at"`
	Type              []string  `json:"type"`
	Flavor            []string  `json:"flavor"`
}

func (app *application) ImageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	imageData, err := app.store.Products.GetProductImage(r.Context(), id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(imageData); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) ProductsHandler(w http.ResponseWriter, r *http.Request) {

	products, err := app.store.Products.GetAllProducts(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, products)
}

func (app *application) SearchProductHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	searchQuery := query.Get("q")

	typeParam := query.Get("type")
	types := []string{}
	if typeParam != "" {
		types = strings.Split(typeParam, ",")
	}

	flavorParam := query.Get("flavor")
	flavors := []string{}
	if flavorParam != "" {
		flavors = strings.Split(flavorParam, ",")
	}

	priceParam := query.Get("price")
	min, max := 0, 0
	if priceParam != "" {

		price, err := strconv.Atoi(priceParam)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		switch price {
		case 1:
			max = 30000
		case 2:
			min = 30000
			max = 60000
		case 3:
			min = 60000
			max = math.MaxInt32
		}
	}

	products, err := app.store.Products.GetProductByFilter(r.Context(), searchQuery, types, flavors, min, max)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, products)
}

func (app *application) SortProductHandler(w http.ResponseWriter, r *http.Request) {
	var sortField string
	var limit int
	var err error

	if topParam := r.URL.Query().Get("top"); topParam != "" {
		sortField = "sold"
		limit, err = strconv.Atoi(topParam)
	} else if newParam := r.URL.Query().Get("new"); newParam != "" {
		sortField = "created_at"
		limit, err = strconv.Atoi(newParam)
	} else {
		app.badRequestResponse(w, r, errors.New("missing 'top' or 'new' query parameter"))
		return
	}

	if err != nil || limit <= 0 {
		app.badRequestResponse(w, r, errors.New("'top' or 'new' must be a positive integer"))
		return
	}

	products, err := app.store.Products.GetBySort(r.Context(), sortField, limit)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, products)
}

func (app *application) DetailProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	products, err := app.store.Products.GetDetailProduct(r.Context(), id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, products)
}

func (app *application) CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	var payload ProductPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Name == "" || payload.Price <= 0 || payload.Stock < 0 {
		app.badRequestResponse(w, r, errors.New("Invalid product data"))
		return
	}

	var imageData []byte

	if payload.Image != "" {
		base64Data := strings.Split(payload.Image, ",")
		if len(base64Data) != 2 {
			app.badRequestResponse(w, r, errors.New("invalid base64 image format"))
			return
		}

		var err error
		imageData, err = base64.StdEncoding.DecodeString(base64Data[1])
		if err != nil {
			app.badRequestResponse(w, r, errors.New("fail to conver image from base64"))
			return
		}
	}

	product := store.Product{
		Name:              payload.Name,
		HeaderDescription: payload.HeaderDescription,
		Description:       payload.Description,
		Price:             payload.Price,
		Stock:             payload.Stock,
		Sold:              payload.Sold,
		Image:             imageData,
		CreatedAt:         time.Now(),
		Type:              payload.Type,
		Flavor:            payload.Flavor,
	}

	productID, err := app.store.Products.CreateProduct(r.Context(), &product)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	response := store.Product{
		ID: productID,
	}

	writeJSON(w, http.StatusOK, response)
}

func (app *application) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var payload map[string]any
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if len(payload) == 0 {
		app.badRequestResponse(w, r, errors.New("no data provided for update"))
		return
	}

	if imageData, ok := payload["image"].(string); ok && imageData != "" {
		base64Data := strings.Split(imageData, ",")
		if len(base64Data) != 2 {
			app.badRequestResponse(w, r, errors.New("invalid base64 image format"))
			return
		}
		
		imageBinary, err := base64.StdEncoding.DecodeString(base64Data[1])
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
		payload["image"] = imageBinary
	}

	product, err := mapToProduct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.store.Products.UpdateProduct(r.Context(), product, id); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}

func (app *application) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := app.store.Products.DeleteProduct(r.Context(), id); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}

func (app *application) DeleteProductInCartHandler(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "id")

	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	if err := app.store.Transaction.DeleteProduct(r.Context(), user.ID, product_id); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)

}

func (app *application) IncQuantityHandler(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "id")

	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	if err := app.store.Transaction.IncrementQuantity(r.Context(), user.ID, product_id); err != nil {
		app.internalServerError(w, r, err)
		return 
	}

	writeJSON(w, http.StatusNoContent, nil)
}

func (app *application) DecQuantityHandler(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "id")

	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	if err := app.store.Transaction.DecrementQuantity(r.Context(), user.ID, product_id); err != nil {
		app.internalServerError(w, r, err)
		return 
	}

	writeJSON(w, http.StatusNoContent, nil)
}