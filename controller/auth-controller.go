package controller

import (
	"net/http"

	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(context *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *authController {
	return &authController{authService: authService}
}

func (controller *authController) Register(context *gin.Context) {
	var userDto dto.User
	// userDto := dto.User{}
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
