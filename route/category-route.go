package route

import (
	"github.com/abdurraufraihan/golang-blog-api/controller"
	"github.com/abdurraufraihan/golang-blog-api/middleware"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoute(db *gorm.DB, categoryRouter *gin.RouterGroup) {
	var (
		jwtService         service.JwtService            = service.NewJwtService()
		categoryRepository repository.CategoryRepo       = repository.NewCategoryRepo(db)
		categoryService    service.CategoryService       = service.NewCategoryService(categoryRepository)
		categoryController controller.CategoryController = controller.NewCategoryController(categoryService)
	)
	categoryRouter.GET("", categoryController.All)
	categoryRouter.POST("", middleware.AuthorizeJWT(jwtService), categoryController.Insert)
	categoryRouter.PUT("/:categoryId", middleware.AuthorizeJWT(jwtService), categoryController.Update)
	categoryRouter.DELETE("/:categoryId", middleware.AuthorizeJWT(jwtService), categoryController.DeleteById)
}
