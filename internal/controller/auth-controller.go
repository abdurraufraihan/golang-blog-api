package controller

import (
	"net/http"

	"github.com/abdurraufraihan/golang-blog-api/internal/dto"
	"github.com/abdurraufraihan/golang-blog-api/internal/service"
	"github.com/abdurraufraihan/golang-blog-api/internal/utils"
	"github.com/abdurraufraihan/golang-blog-api/pkg/logger"
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
	logger      *logger.Logger
}

func NewAuthController(
	authService service.AuthService,
	jwtService service.JwtService,
	logger *logger.Logger,
) *authController {
	return &authController{
		authService: authService, jwtService: jwtService, logger: logger,
	}
}

// Login             godoc
// @Summary      User login
// @Description  Responds with token pair as JSON.
// @Tags         auth
// @Produce      json
// @Param data body dto.Login true "Login dto"
// @Success      200
// @Router       /login [post]
func (controller *authController) Login(context *gin.Context) {
	var loginDto dto.Login
	err := context.ShouldBindJSON(&loginDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	isValidCredential, userId := controller.authService.VerifyCredential(
		loginDto.Email, loginDto.Password)
	if isValidCredential {
		tokenPair := controller.jwtService.GenerateTokenPair(userId)
		context.JSON(http.StatusOK, utils.GetResponse(tokenPair))
		return
	}
	context.JSON(
		http.StatusBadRequest, utils.GetErrorResponse("invalid credential"))
	controller.logger.Error().Msg("invalid credential")
}

// Register             godoc
// @Summary      User register
// @Description  Responds with user data as JSON.
// @Tags         auth
// @Produce      json
// @Param data body dto.User true "User dto"
// @Success      200
// @Router       /signup [post]
func (controller *authController) Register(context *gin.Context) {
	var userDto dto.User
	err := context.ShouldBindJSON(&userDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	result, user := controller.authService.Register(userDto)
	if result.Error != nil {
		context.JSON(
			http.StatusBadRequest, utils.GetErrorResponse(result.Error.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	context.JSON(http.StatusOK, utils.GetResponse(user))
}

// Verify Token             godoc
// @Summary      Verify user token
// @Description  Responds with is_valid status as JSON.
// @Tags         auth
// @Produce      json
// @Param data body dto.Token true "Token dto"
// @Success      200
// @Router       /token/verify [post]
func (controller *authController) VerifyToken(context *gin.Context) {
	tokenDto := dto.Token{}
	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	token, _ := utils.ValidateToken(tokenDto.Token)
	if token == nil || !token.Valid {
		context.AbortWithStatusJSON(
			http.StatusBadRequest, utils.GetErrorResponse("Invalid Token"))
		controller.logger.Error().Msg("Invalid Token")
		return
	}
	context.JSON(http.StatusOK, utils.GetResponse(gin.H{"is_valid": true}))
}

// Refresh Token             godoc
// @Summary      Refresh user token
// @Description  Responds with token pair as JSON.
// @Tags         auth
// @Produce      json
// @Param data body dto.Token true "Token dto"
// @Success      200
// @Router       /token/refresh [post]
func (controller *authController) RefreshToken(context *gin.Context) {
	tokenDto := dto.Token{}
	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	token, err := utils.ValidateToken(tokenDto.Token)
	if token == nil || !token.Valid {
		context.AbortWithStatusJSON(
			http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		context.JSON(
			http.StatusOK, controller.jwtService.GenerateTokenPair(claims["user_id"]))
	} else {
		context.AbortWithStatusJSON(
			http.StatusBadRequest, utils.GetErrorResponse("Failed to claim token"))
		controller.logger.Error().Msg("Failed to claim token")
	}
}
