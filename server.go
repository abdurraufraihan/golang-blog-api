package main

import (
	"github.com/abdurraufraihan/golang-blog-api/config"
	"github.com/abdurraufraihan/golang-blog-api/controller"
	"github.com/abdurraufraihan/golang-blog-api/middleware"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.ConnectWithDB()
	postRepository repository.PostRepo       = repository.NewPostRepo(db)
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)

	categoryRepository repository.CategoryRepo       = repository.NewCategoryRepo(db)
	categoryService    service.CategoryService       = service.NewCategoryService(categoryRepository)
	categoryController controller.CategoryController = controller.NewCategoryController(categoryService)

	jwtService     service.JwtService        = service.NewJwtService()
	authRepository repository.AuthRepo       = repository.NewAuthRepo(db)
	authService    service.AuthService       = service.NewAuthService(authRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDbConnection(db)
	router := gin.Default()

	router.GET("/posts", postController.All)
	router.GET("/posts/:id", postController.FindById)
	router.POST("/posts", middleware.AuthorizeJWT(jwtService), postController.Insert)

	router.GET("/categories", categoryController.All)

	router.POST("/signup", authController.Register)
	router.POST("/login", authController.Login)

	router.Run("localhost:8080")
}
