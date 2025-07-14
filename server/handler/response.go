package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func WrongContentType(w http.ResponseWriter, r *http.Request, contentType string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ResponseError{
		Message: fmt.Sprintf("Content-Type must be %s", contentType),
		Status:  http.StatusBadRequest,
	})
}

func BadRequest(w http.ResponseWriter, r *http.Request, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ResponseError{
		Message: message,
		Status:  http.StatusBadRequest,
	})
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(ResponseError{
		Message: "Unauthorized",
		Status:  http.StatusUnauthorized,
	})
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ResponseError{
		Message: "Internal Server Error",
		Status:  http.StatusInternalServerError,
	})
}
