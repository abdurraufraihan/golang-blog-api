package repository

import (
	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(user model.User) (*gorm.DB, model.User)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepo {
	return &authRepo{db: db}
}

func (repo *authRepo) Register(user model.User) (*gorm.DB, model.User) {
	userResult := repo.db.Create(&user)
	return userResult, user
}
