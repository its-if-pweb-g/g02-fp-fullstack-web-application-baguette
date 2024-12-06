package main

import (
	"api/internal/store"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=256"`
	Email    string `json:"email" validate:"required,email,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
	Phone    string `json:"phone"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

type RegisterAndLoginResponse struct {
	Token string `json:"token"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &store.User{
		Name:  payload.Username,
		Email: payload.Email,
		Phone: payload.Phone,
		Role:  "user",
		Addrres: "",
	}

	if err := user.SetPassword(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	resultID, err := app.store.Users.Create(r.Context(), user)
	if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.badRequestResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	claims := jwt.MapClaims{
		"ID":  resultID,
		"role": "user",
		"exp": time.Now().Add(app.config.auth.exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.iss,
		"aud": app.config.auth.iss,
	}

	token, err := app.authenticator.GenerateToken(claims)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	reponse := RegisterAndLoginResponse{Token: token}
	if err := writeJSON(w, http.StatusOK, reponse); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	result, err := app.store.Users.GetByEmail(r.Context(),payload.Email) 
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := result.Compare(payload.Password); err != nil {
		app.unauthorizedErrorResponse(w, r, err)
	}

	claims := jwt.MapClaims{
		"ID":  result.ID,
		"role": result.Role,
		"exp": time.Now().Add(app.config.auth.exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.iss,
		"aud": app.config.auth.iss,
	}

	token, err := app.authenticator.GenerateToken(claims)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	reponse := RegisterAndLoginResponse{Token: token}
	if err := writeJSON(w, http.StatusOK, reponse); err != nil {
		app.internalServerError(w, r, err)
	}
}
