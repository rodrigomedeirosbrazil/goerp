package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken() (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwt_key := os.Getenv("JWT_KEY")

	t, err := token.SignedString([]byte(jwt_key))
	if err != nil {
		return "", err
	}

	return t, nil
}
