// Package utils provides helper functions for HTTP response formatting.
package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithJSON writes a JSON response with the given status code and payload.
func RespondWithJSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to write JSON response: %v", err)
	}
}

// BasicResponse represents a standard JSON structure for success/error messages.
type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// RespondWithSuccess sends a success message as JSON with the specified status code.
func RespondWithSuccess(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(BasicResponse{Success: true, Message: message}); err != nil {
		log.Printf("failed to write success response: %v", err)
	}
}

// RespondWithError sends an error message as JSON with the specified status code.
func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if encodeErr := json.NewEncoder(w).Encode(BasicResponse{Success: false, Message: err.Error()}); encodeErr != nil {
		log.Printf("failed to write error response: %v", encodeErr)
	}
}
