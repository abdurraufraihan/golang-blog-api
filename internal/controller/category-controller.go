package controller

import (
	"net/http"
	"strconv"

	"github.com/abdurraufraihan/golang-blog-api/internal/dto"
	"github.com/abdurraufraihan/golang-blog-api/internal/serializer"
	"github.com/abdurraufraihan/golang-blog-api/internal/service"
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

// GetCategories             godoc
// @Summary      Get categories list
// @Description  Responds with the list of all categories as JSON.
// @Tags         categories
// @Produce      json
// @Success      200  {object}  serializer.CategoryResponse
// @Router       /categories [get]
func (controller *categoryController) All(context *gin.Context) {
	categories := controller.categoryService.All()
	serializer := serializer.CategoriesSerializer{Categories: categories}
	context.JSON(http.StatusOK, serializer.Response())
}

// InsertCategory             godoc
// @Summary      Insert category
// @Description  Responds with category as JSON.
// @Tags         categories
// @Produce      json
// @Param data body dto.Category true "Category dto"
// @Success      201  {object}  serializer.CategoryResponse
// @Router       /categories [post]
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

// UpdateCategory             godoc
// @Summary      Update category
// @Description  Responds with category as JSON.
// @Tags         categories
// @Produce      json
// @Param        id  path      uint  true  "update category by id"
// @Param data body dto.Category true "Category dto"
// @Success      200  {object}  serializer.CategoryResponse
// @Router       /categories/{id} [put]
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

// DeleteCategory             godoc
// @Summary      Delete category
// @Description  Responds with category as JSON.
// @Tags         categories
// @Produce      json
// @Param        id  path      uint  true  "delete category by id"
// @Success      204
// @Router       /categories/{id} [delete]
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
