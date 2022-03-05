package service

import (
	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/model"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/mashingan/smapping"
)

type PostService interface {
	All() []model.Post
	FindById(postId uint64) (model.Post, error)
	Insert(postDto dto.Post) model.Post
	Update(postId uint64, postDto dto.Post) (model.Post, error)
}

type postService struct {
	postRepo repository.PostRepo
}

func NewPostService(postRepo repository.PostRepo) *postService {
	return &postService{
		postRepo: postRepo,
	}
}

func (service *postService) All() []model.Post {
	return service.postRepo.AllPost()
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
	fillErr := smapping.FillStruct(&post, smapping.MapFields(&postDto))
	if fillErr != nil {
		panic(fillErr)
	}
	service.postRepo.Save(&post)
	postWithCategory, err := service.postRepo.FindByIdWithCategory(postId)
	if err != nil {
		return postWithCategory, err
	}
	return postWithCategory, nil
}
