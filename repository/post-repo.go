package repository

import (
	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/gorm"
)

type PostRepo interface {
	Insert(post model.Post) model.Post
	AllPost() []model.Post
	FindByPostId(postId uint64) (model.Post, error)
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
	repo.db.Preload("Category").Find(&posts)
	return posts
}

func (repo *postRepo) FindByPostId(postId uint64) (model.Post, error) {
	var post model.Post
	result := repo.db.Preload("Category").First(&post, postId)
	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}
