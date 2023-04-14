package route

import (
	"github.com/abdurraufraihan/golang-blog-api/internal/controller"
	"github.com/abdurraufraihan/golang-blog-api/internal/middleware"
	"github.com/abdurraufraihan/golang-blog-api/internal/repository"
	"github.com/abdurraufraihan/golang-blog-api/internal/service"
	"github.com/abdurraufraihan/golang-blog-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoute(db *gorm.DB, categoryRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		categoryRepository repository.CategoryRepo = repository.NewCategoryRepo(db)
		categoryService    service.CategoryService = service.
					NewCategoryService(categoryRepository)
		categoryController controller.CategoryController = controller.
					NewCategoryController(categoryService, logger)
	)
	categoryRouter.GET("", categoryController.All)
	categoryRouter.POST("", middleware.AuthorizeJWT(), categoryController.Insert)
	categoryRouter.PUT(
		"/:categoryId", middleware.AuthorizeJWT(), categoryController.Update)
	categoryRouter.DELETE(
		"/:categoryId", middleware.AuthorizeJWT(), categoryController.DeleteById)
}
