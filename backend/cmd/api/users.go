package main

import (
	"api/internal/store"
	"errors"
	"net/http"

)

type UpdateUserPayload struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Role     string `json:"role,omitempty"`
	Addrres  string `json:"address,omitempty"`
}

func (app *application) UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	user.Password = []byte("")

	writeJSON(w, http.StatusOK, user)
}

func (app *application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.unauthorizedErrorResponse(w, r, errors.New("you are not registered"))
		return
	}

	var payload store.User
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.store.Users.Update(r.Context(), &payload, user.ID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
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

	_, cart, err := app.store.Transaction.GetUserCart(r.Context(), user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}


	writeJSON(w, http.StatusOK, cart)
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

	if err := app.store.Transaction.AddProduct(r.Context(), payload, user.ID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

