package service

import (
	"github.com/abdurraufraihan/golang-blog-api/internal/dto"
	"github.com/abdurraufraihan/golang-blog-api/internal/model"
	"github.com/abdurraufraihan/golang-blog-api/internal/repository"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type PostService interface {
	All(limit string, offset string) (int64, []model.Post)
	FindById(postId uint64) (model.Post, error)
	Insert(postDto dto.Post) model.Post
	Update(postId uint64, postDto dto.Post) (model.Post, error)
	DeleteById(postId uint64) *gorm.DB
}

type postService struct {
	postRepo repository.PostRepo
}

func NewPostService(postRepo repository.PostRepo) *postService {
	return &postService{
		postRepo: postRepo,
	}
}

func (service *postService) All(limit string, offset string) (int64, []model.Post) {
	return service.postRepo.PostCount(), service.postRepo.AllPost(limit, offset)
}

func (service *postService) FindById(postId uint64) (model.Post, error) {
	return service.postRepo.FindByIdWithCategory(postId)
}

func (service *postService) Insert(postDto dto.Post) model.Post {
	postModel := model.Post{}
	err := smapping.FillStruct(&postModel, smapping.MapFields(&postDto))
	if err != nil {
		panic(err)
	}
	res := service.postRepo.Insert(postModel)
	return res
}

func (service *postService) Update(
	postId uint64, postDto dto.Post,
) (model.Post, error) {
	post, err := service.postRepo.FindById(postId)
	if err != nil {
		return post, err
	}
	if postDto.Image == "" {
		postDto.Image = post.Image
	}
	if fillErr := smapping.FillStruct(&post, smapping.MapFields(&postDto)); fillErr != nil {
		panic(fillErr)
	}
	if result := service.postRepo.Save(&post); result.Error != nil {
		return post, result.Error
	}
	postWithCategory, err := service.postRepo.FindByIdWithCategory(postId)
	if err != nil {
		return postWithCategory, err
	}
	return postWithCategory, nil
}

func (service *postService) DeleteById(postId uint64) *gorm.DB {
	return service.postRepo.DeleteById(postId)
}
