// Package healthcheck provides a simple HTTP handler to report service status.
package healthcheck

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler responds with a JSON status to indicate the API is running.
func Handler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "API is up and running",
	}); err != nil {
		log.Printf("failed to write health check response: %v", err)
	}
}
