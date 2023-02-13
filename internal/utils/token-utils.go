package utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GetSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey != "" {
		secretKey = "fc93cb07e1ad92898527100e58a1cf1d1e7"
	}
	return secretKey
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetSecretKey()), nil
	})
}
