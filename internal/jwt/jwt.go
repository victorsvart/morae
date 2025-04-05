// Package jwt provides utilities for generating JSON Web Tokens for authentication.
package jwt

import (
	"fmt"
	"os"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte(os.Getenv("SECRET_KEY"))

// GenerateJWT creates and returns a signed JWT for the given email address.
// The token expires after 24 hours.
func GenerateJWT(emailaddress string) (string, error) {
	if len(secretkey) == 0 {
		return "", fmt.Errorf("secret key is not set")
	}

	claims := jwtlib.MapClaims{
		"username": emailaddress,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString(secretkey)
}
