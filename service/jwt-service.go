package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateTokenPair(userId uint64) map[string]string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	issuer    string
	secretKey string
}

func NewJwtService() *jwtService {
	return &jwtService{issuer: "raihan", secretKey: getSecretKey()}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "d47s4654dsf545sdf54sdf8"
	}
	return secretKey
}

func (service *jwtService) GenerateTokenPair(userId uint64) map[string]string {
	claims := &jwtCustomClaim{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 15).Unix(), // 15 days
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(), // 1 month
		Issuer:    service.issuer,
		IssuedAt:  time.Now().Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return map[string]string{"access_token": tokenString, "refresh_token": refreshTokenString}
}

func (service *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
