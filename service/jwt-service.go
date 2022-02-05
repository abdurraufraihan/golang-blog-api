package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(userId uint64) string
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
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
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
