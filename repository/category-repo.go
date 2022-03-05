package repository

import (
	"errors"

	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	AllCategories() []model.Category
	Insert(category model.Category) model.Category
	GetById(id uint64) (model.Category, error)
	Save(category *model.Category)
	DeleteById(categoryId uint64) *gorm.DB
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepo {
	return &categoryRepo{db: db}
}

func (repo *categoryRepo) AllCategories() []model.Category {
	category := []model.Category{}
	repo.db.Order("id desc").Find(&category)
	return category
}

func (repo *categoryRepo) Insert(category model.Category) model.Category {
	repo.db.Create(&category)
	return category
}

func (repo *categoryRepo) GetById(id uint64) (model.Category, error) {
	category := model.Category{}
	err := repo.db.First(&category, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return category, err
	}
	return category, nil
}

func (repo *categoryRepo) Save(category *model.Category) {
	repo.db.Save(category)
}

func (repo *categoryRepo) DeleteById(categoryId uint64) *gorm.DB {
	return repo.db.Delete(&model.Category{}, categoryId)
}
