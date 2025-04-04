package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(emailaddress string) (string, error) {
	// Ensure the secret key is not empty
	if len(secretkey) == 0 {
		return "", fmt.Errorf("secret key is not set")
	}

	claims := jwt.MapClaims{
		"username": emailaddress,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// Use a valid HMAC signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretkey)
}
