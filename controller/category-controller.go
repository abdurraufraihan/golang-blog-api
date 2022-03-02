package controller

import (
	"net/http"

	"github.com/abdurraufraihan/golang-blog-api/serializer"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	All(context *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(
	categoryService service.CategoryService,
) *categoryController {
	return &categoryController{categoryService: categoryService}
}

func (controller *categoryController) All(context *gin.Context) {
	categories := controller.categoryService.All()
	serializer := serializer.CategoriesSerializer{Categories: categories}
	context.JSON(http.StatusOK, serializer.Response())
}
