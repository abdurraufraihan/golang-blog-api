package repository

import (
	"github.com/abdurraufraihan/golang-blog-api/internal/model"
	"gorm.io/gorm"
)

type CommentRepo interface {
	Insert(comment model.Comment) model.Comment
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *commentRepo {
	return &commentRepo{db: db}
}

func (repo *commentRepo) Insert(comment model.Comment) model.Comment {
	repo.db.Create(&comment)
	repo.db.First(&comment, comment.ID)
	return comment
}
