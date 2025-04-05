package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(emailaddress string) (string, error) {
	if len(secretkey) == 0 {
		return "", fmt.Errorf("secret key is not set")
	}

	claims := jwt.MapClaims{
		"username": emailaddress,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretkey)
}
