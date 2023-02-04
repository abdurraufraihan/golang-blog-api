package service

import (
	"github.com/abdurraufraihan/golang-blog-api/internal/dto"
	"github.com/abdurraufraihan/golang-blog-api/internal/model"
	"github.com/abdurraufraihan/golang-blog-api/internal/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(userDto dto.User) (*gorm.DB, model.User)
	VerifyCredential(email string, passsword string) (bool, uint64)
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

func (service *authService) VerifyCredential(email string, passsword string) (bool, uint64) {
	user := service.authRepo.FindByEmail(email)
	if (user != model.User{}) {
		return comparePassword([]byte(user.Password), []byte(passsword)), user.ID
	}
	return false, 0
}

func comparePassword(hashedPass []byte, plainPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPass, plainPass)
	if err != nil {
		return false
	}
	return true
}
