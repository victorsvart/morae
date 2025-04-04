package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func RespondWithSuccess(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(BasicResponse{Success: true, Message: message})
}

func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(BasicResponse{Success: false, Message: err.Error()})
}
