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

func (app *application) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ngrok-skip-browser-warning")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
			fmt.Print(cookie)
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
