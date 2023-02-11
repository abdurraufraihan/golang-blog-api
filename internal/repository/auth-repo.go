package repository

import (
	"github.com/abdurraufraihan/golang-blog-api/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(user model.User) (*gorm.DB, model.User)
	FindByEmail(email string) (*gorm.DB, model.User)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepo {
	return &authRepo{db: db}
}

func (repo *authRepo) Register(user model.User) (*gorm.DB, model.User) {
	user.Password = hashAndSalt([]byte(user.Password))
	userResult := repo.db.Create(&user)
	return userResult, user
}

func (repo *authRepo) FindByEmail(email string) (*gorm.DB, model.User) {
	user := model.User{}
	userResult := repo.db.Where("email = ?", email).Take(&user)
	return userResult, user
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash password")
	}
	return string(hash)
}
