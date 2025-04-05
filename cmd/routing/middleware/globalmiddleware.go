// Package middleware provides reusable HTTP middleware for logging, JSON headers, and authentication.
package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the HTTP method, URL path, and duration of each request.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s at %v", r.Method, r.URL.Path, start)
		next(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// JSONMiddleware sets the Content-Type of the response to application/json.
func JSONMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
