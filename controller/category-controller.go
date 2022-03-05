package controller

import (
	"net/http"
	"strconv"

	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/serializer"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	DeleteById(context *gin.Context)
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

func (controller *categoryController) Insert(context *gin.Context) {
	categoryDto := dto.Category{}
	err := context.ShouldBindJSON(&categoryDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	category := controller.categoryService.Insert(categoryDto)
	serializer := serializer.CategorySerializer{Category: category}
	context.JSON(http.StatusOK, serializer.Response())
}

func (controller *categoryController) Update(context *gin.Context) {
	categoryDto := dto.Category{}
	err := context.ShouldBindJSON(&categoryDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	categoryId, _ := strconv.ParseUint(context.Param("categoryId"), 10, 64)
	category, err := controller.categoryService.Update(categoryId, categoryDto)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	serializer := serializer.CategorySerializer{Category: category}
	context.JSON(http.StatusOK, serializer.Response())
}

func (controller *categoryController) DeleteById(context *gin.Context) {
	categoryId, _ := strconv.ParseUint(context.Param("categoryId"), 10, 64)
	result := controller.categoryService.DeleteById(categoryId)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	} else if result.RowsAffected < 1 {
		context.JSON(http.StatusNotFound, gin.H{"error": "category does not exists"})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}
