package service

import (
	"github.com/abdurraufraihan/golang-blog-api/model"
	"github.com/abdurraufraihan/golang-blog-api/repository"
)

type CategoryService interface {
	All() []model.Category
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
