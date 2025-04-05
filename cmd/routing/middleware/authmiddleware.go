// Package middleware provides reusable HTTP middleware for handling logging, JSON responses, and authentication.
package middleware

import (
	"fmt"
	"morae/cmd/config"
	"net/http"
	"os"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

// verifyJWT parses and verifies a JWT token using the HS256 algorithm and the secret key from environment variables.
func verifyJWT(tokenString string) error {
	_, err := jwtlib.Parse(tokenString, func(token *jwtlib.Token) (any, error) {
		if token.Method != jwtlib.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}

		secret := os.Getenv("SECRET_KEY")
		if secret == "" {
			return nil, fmt.Errorf("SECRET_KEY is not set")
		}

		return []byte(secret), nil
	})

	return err
}

// AuthMiddleware is an HTTP middleware that verifies a user's JWT token from cookies.
// It blocks the request with a 401 Unauthorized response if the token is missing or invalid.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(config.GetEnv("AUTH_TOKEN_NAME", "dev_token"))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if err := verifyJWT(cookie.Value); err != nil {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	})
}
