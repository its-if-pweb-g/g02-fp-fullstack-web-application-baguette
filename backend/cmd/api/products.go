package main

import (
	"api/internal/store"
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)



func (app *application) ImageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	imageData, err := app.store.Products.GetProductImage(r.Context(), id);
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

	if err := app.jsonResponse(w, http.StatusOK, products); err != nil {
		app.internalServerError(w, r, err)
	}
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
		if  err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		switch (price) {
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

	if err := app.jsonResponse(w, http.StatusOK, products); err != nil {
		app.internalServerError(w, r, err)
	}
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

	if err := app.jsonResponse(w, http.StatusOK, products); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) DetailProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	products, err := app.store.Products.GetDetailProduct(r.Context(), id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, products); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	
	var payload store.Product
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Name == "" || payload.Price <= 0 || payload.Stock < 0 {
		app.badRequestResponse(w, r, errors.New("Invalid product data"))
		return
	}

	productID, err := app.store.Products.CreateProduct(r.Context(), &payload)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	response := store.Product{
		ID: productID,
	}

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	
	id := chi.URLParam(r, "id")
	
	var payload store.Product
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.store.Products.UpdateProduct(r.Context(), &payload, id); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := app.store.Products.DeleteProduct(r.Context(), id); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

