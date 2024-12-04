package main

import (
	"api/internal/store"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const userCtx string = "user"

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("auth token is missing"))

			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("auth header have different type"))
			return
		}

		tokenHeader := parts[1]

		jwtToken, err := app.authenticator.ValidateToken(tokenHeader)

		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		user, err := app.store.Users.GetByID(r.Context(), claims["ID"].(string))
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func (app *application) AdminRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value(userCtx).(*store.User)

		if user.Role != "admin" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("Unoutorized role"))
			return
		}

		next.ServeHTTP(w, r)
	})
}


