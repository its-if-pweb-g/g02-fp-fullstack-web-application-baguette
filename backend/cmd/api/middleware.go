package main

import (
	"api/internal/store"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type userKey string

const userCtx userKey = "user"

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenHeader string

		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenHeader = parts[1]
			}
		}

		if tokenHeader == "" {
			cookie, err := r.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					app.unauthorizedErrorResponse(w, r, fmt.Errorf("auth token is missing"))
					return
				}
				app.internalServerError(w, r, err)
				return
			}
			tokenHeader = cookie.Value
		}

		jwtToken, err := app.authenticator.ValidateToken(tokenHeader)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		userID, ok := claims["ID"].(string)
		if !ok {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("invalid token claims"))
			return
		}

		user, err := app.store.Users.GetByID(r.Context(), userID)
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
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("You are not an admin"))
			return
		}

		next.ServeHTTP(w, r)
	})
}


