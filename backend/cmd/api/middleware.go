package main

import (
	// "fmt"
	"net/http"
	// "strconv"
	// "strings"

	// "github.com/golang-jwt/jwt/v5"
)

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		// authHeader := r.Header.Get("Authorization")

		// if authHeader == "" {
		// 	app.unauthorizedErrorResponse(w, r, fmt.Errorf("auth token is missing"))

		// 	return
		// }

		// parts := strings.Split(authHeader, " ")

		// if len(parts) != 2 || parts[0] != "Bearer" {
		// 	app.unauthorizedErrorResponse(w, r, fmt.Errorf("auth header have different type"))

		// 	return
		// }

		// token := parts[1]

		// jwtToken, err := app.authenticator.ValidateToken(token)

		// if err != nil {
		// 	app.unauthorizedErrorResponse(w, r, err)
		// 	return
		// }

		// claims, _ := jwtToken.Claims.(jwt.MapClaims)

		// userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user"]), 10, 64)

		// if err != nil {
		// 	app.unauthorizedErrorResponse(w, r, err)
		// 	return
		// }

		// ctx := r.Context()

	})
}


func (app *application) AdminRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
