package service

import (
	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/model"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/mashingan/smapping"
)

type PostService interface {
	Insert(postDto dto.Post) model.Post
	All() []model.Post
}

type postService struct {
	postRepo repository.PostRepo
}

func NewPostService(postRepo repository.PostRepo) *postService {
	return &postService{
		postRepo: postRepo,
	}
}

func (service postService) Insert(postDto dto.Post) model.Post {
	postModel := model.Post{}
	err := smapping.FillStruct(&postModel, smapping.MapFields(&postDto))
	if err != nil {
		panic(err)
	}
	res := service.postRepo.Insert(postModel)
	return res
}

func (service postService) All() []model.Post {
	return service.postRepo.AllPost()
}
