package main

import (
	"api/internal/store"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &envelope{Error: message})
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}

	return writeJSON(w, status, &envelope{Data: data})
}

func mapToProduct(payload map[string]any) (*store.Product, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var product store.Product
	err = json.Unmarshal(jsonData, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func mapToUser(payload map[string]any) (*store.User, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var user store.User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
