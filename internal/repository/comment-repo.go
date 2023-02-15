package repository

import (
	"strconv"

	"github.com/abdurraufraihan/golang-blog-api/internal/model"
	"gorm.io/gorm"
)

type CommentRepo interface {
	AllCommentByPostId(limit string, offset string, postId uint) []model.Comment
	Insert(comment model.Comment) model.Comment
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *commentRepo {
	return &commentRepo{db: db}
}

func (repo *commentRepo) AllCommentByPostId(
	limit string, offset string, postId uint,
) []model.Comment {
	commentLimit, err := strconv.Atoi(limit)
	if err != nil || commentLimit == 0 {
		commentLimit = -1
	}
	commentOffset, err := strconv.Atoi(offset)
	if err != nil || commentOffset == 0 {
		commentOffset = -1
	}
	var comments []model.Comment
	repo.db.Where("post_id = ?", postId).Limit(commentLimit).
		Offset(commentOffset).Find(&comments)
	return comments
}

func (repo *commentRepo) Insert(comment model.Comment) model.Comment {
	repo.db.Create(&comment)
	repo.db.First(&comment, comment.ID)
	return comment
}
