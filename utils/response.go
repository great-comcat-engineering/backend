package utils

import (
	"encoding/json"
	"greatcomcatengineering.com/backend/models"
	"net/http"
)

// RespondWithJSON sends a JSON response with a generic content type.
func RespondWithJSON[T any](w http.ResponseWriter, status int, message string, content T) {
	response := models.ApiResponse[T]{
		Status:  status,
		Message: message,
		Content: content,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// RespondWithError is a convenience wrapper for sending error messages using the ApiResponse structure.
func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON[interface{}](w, status, message, nil)
}
