package service

import (
	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/model"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/mashingan/smapping"
)

type CategoryService interface {
	All() []model.Category
	Insert(categoryDto dto.Category) model.Category
}

type categoryService struct {
	categoryRepo repository.CategoryRepo
}

func NewCategoryService(categoryRepo repository.CategoryRepo) *categoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (service *categoryService) All() []model.Category {
	return service.categoryRepo.AllCategories()
}

func (service *categoryService) Insert(categoryDto dto.Category) model.Category {
	categoryModel := model.Category{}
	err := smapping.FillStruct(&categoryModel, smapping.MapFields(&categoryDto))
	if err != nil {
		panic(err)
	}
	return service.categoryRepo.Insert(categoryModel)
}
