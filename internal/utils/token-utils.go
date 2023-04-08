package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
			return nil, fmt.Errorf(
				"unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetSecretKey()), nil
	})
}

func GetTokenString(context *gin.Context) string {
	bearerToken := context.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GetUserIDFromToken(tokenString string) (uint, error) {
	token, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
