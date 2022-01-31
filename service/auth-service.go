package service

import (
	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/model"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(userDto dto.User) (*gorm.DB, model.User)
}

type authService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(authRepo repository.AuthRepo) *authService {
	return &authService{authRepo: authRepo}
}

func (service *authService) Register(userDto dto.User) (*gorm.DB, model.User) {
	userModel := model.User{}
	err := smapping.FillStruct(&userModel, smapping.MapFields(&userDto))
	if err != nil {
		panic(err)
	}
	return service.authRepo.Register(userModel)
}
