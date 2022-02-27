package repository

import (
	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	AllCategories() []model.Category
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepo {
	return &categoryRepo{db: db}
}

func (repo *categoryRepo) AllCategories() []model.Category {
	category := []model.Category{}
	repo.db.Find(&category)
	return category
}
