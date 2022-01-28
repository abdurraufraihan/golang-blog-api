package repository

import (
	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/gorm"
)

type PostRepo interface {
	Insert(post model.Post) model.Post
	AllPost() []model.Post
	FindByPostId(postId uint64) model.Post
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *postRepo {
	return &postRepo{db: db}
}

func (repo *postRepo) Insert(post model.Post) model.Post {
	repo.db.Create(&post)
	return post
}

func (repo *postRepo) AllPost() []model.Post {
	var posts []model.Post
	repo.db.Find(&posts)
	return posts
}

func (repo *postRepo) FindByPostId(postId uint64) model.Post {
	var post model.Post
	repo.db.Find(&post, postId)
	return post
}
