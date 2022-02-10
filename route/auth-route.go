package route

import (
	"github.com/abdurraufraihan/golang-blog-api/controller"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(db *gorm.DB, authRouter *gin.RouterGroup) {
	var (
		jwtService     service.JwtService        = service.NewJwtService()
		authRepository repository.AuthRepo       = repository.NewAuthRepo(db)
		authService    service.AuthService       = service.NewAuthService(authRepository)
		authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	)
	authRouter.POST("/login", authController.Login)
	authRouter.POST("/signup", authController.Register)
}
