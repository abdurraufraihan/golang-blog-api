package controller

import (
	"net/http"

	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	VerifyToken(context *gin.Context)
	RefreshToken(context *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JwtService
}

func NewAuthController(
	authService service.AuthService, jwtService service.JwtService,
) *authController {
	return &authController{authService: authService, jwtService: jwtService}
}

func (controller *authController) Login(context *gin.Context) {
	var loginDto dto.Login
	err := context.ShouldBindJSON(&loginDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValidCredential, userId :=
		controller.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if isValidCredential {
		tokenPair := controller.jwtService.GenerateTokenPair(userId)
		context.JSON(http.StatusOK, tokenPair)
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{"error": "invalid credential"})
}

func (controller *authController) Register(context *gin.Context) {
	var userDto dto.User
	err := context.ShouldBindJSON(&userDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, user := controller.authService.Register(userDto)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, result.Error)
		return
	}
	context.JSON(http.StatusOK, user)
}

func (controller *authController) VerifyToken(context *gin.Context) {
	tokenDto := dto.Token{}
	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, _ := controller.jwtService.ValidateToken(tokenDto.Token)
	if token == nil || !token.Valid {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"is_valid": false})
		return
	}
	context.JSON(http.StatusOK, gin.H{"is_valid": true})
}

func (controller *authController) RefreshToken(context *gin.Context) {
	tokenDto := dto.Token{}
	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := controller.jwtService.ValidateToken(tokenDto.Token)
	if token == nil || !token.Valid {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		context.JSON(http.StatusOK, controller.jwtService.GenerateTokenPair(claims["user_id"]))
	} else {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "can not able to claim"})
	}
}
