package middleware

import (
	"net/http"

	"github.com/abdurraufraihan/golang-blog-api/internal/service"
	"github.com/abdurraufraihan/golang-blog-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{"error": "No token found"},
			)
			return
		}
		tokenString := authHeader[7:] // remove Bearer from token
		token, err := utils.ValidateToken(tokenString)
		if token == nil || !token.Valid {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{"error": err.Error()},
			)
			return
		}
	}
}
