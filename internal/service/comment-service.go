package service

import (
	"github.com/abdurraufraihan/golang-blog-api/internal/dto"
	"github.com/abdurraufraihan/golang-blog-api/internal/model"
	"github.com/abdurraufraihan/golang-blog-api/internal/repository"
	"github.com/mashingan/smapping"
)

type CommentService interface {
	Insert(commentDto dto.Comment, postId, userId uint) model.Comment
}

type commentService struct {
	commentRepo repository.CommentRepo
}

func NewCommentService(commentRepo repository.CommentRepo) *commentService {
	return &commentService{commentRepo: commentRepo}
}

func (service *commentService) Insert(commentDto dto.Comment, postId, userId uint) model.Comment {
	commentModel := model.Comment{}
	err := smapping.FillStruct(&commentModel, smapping.MapFields(&commentDto))
	if err != nil {
		panic(err)
	}
	commentModel.PostID = postId
	commentModel.UserID = userId
	return service.commentRepo.Insert(commentModel)
}
