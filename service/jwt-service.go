package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(userId uint64) string
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

func (service *jwtService) GenerateToken(userId uint64) string {
	claims := &jwtCustomClaim{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return ss
}

func (service *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
