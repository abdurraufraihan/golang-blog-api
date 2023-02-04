package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateTokenPair(userId interface{}) map[string]string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserId interface{} `json:"user_id"`
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
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey != "" {
		secretKey = "fc93cb07e1ad92898527100e58a1cf1d1e7"
	}
	return secretKey
}

func (service *jwtService) getTokenClaims(
	userId interface{}, expiryDays int,
) *jwtCustomClaim {
	return &jwtCustomClaim{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, expiryDays).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
}

func (service *jwtService) GenerateTokenPair(userId interface{}) map[string]string {
	tokenClaims := service.getTokenClaims(userId, 15)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	refreshTokenClaims := service.getTokenClaims(userId, 30)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return map[string]string{"access_token": tokenString, "refresh_token": refreshTokenString}
}

func (service *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
