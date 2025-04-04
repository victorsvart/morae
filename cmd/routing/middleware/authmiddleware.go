package middleware

import (
	"fmt"
	"morae/cmd/config"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func verifyJWT(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
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
