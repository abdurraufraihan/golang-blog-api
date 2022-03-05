package route

import (
	"github.com/abdurraufraihan/golang-blog-api/controller"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoute(db *gorm.DB, categoryRouter *gin.RouterGroup) {
	var (
		categoryRepository repository.CategoryRepo       = repository.NewCategoryRepo(db)
		categoryService    service.CategoryService       = service.NewCategoryService(categoryRepository)
		categoryController controller.CategoryController = controller.NewCategoryController(categoryService)
	)
	categoryRouter.GET("", categoryController.All)
	categoryRouter.POST("", categoryController.Insert)
	categoryRouter.PUT("/:categoryId", categoryController.Update)
	categoryRouter.DELETE("/:categoryId", categoryController.DeleteById)
}
