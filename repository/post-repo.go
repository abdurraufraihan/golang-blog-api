package repository

import (
	"errors"

	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/gorm"
)

type PostRepo interface {
	AllPost() []model.Post
	FindByIdWithCategory(postId uint64) (model.Post, error)
	FindById(postId uint64) (model.Post, error)
	Insert(post model.Post) model.Post
	Save(post *model.Post) *gorm.DB
	DeleteById(postId uint64) *gorm.DB
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *postRepo {
	return &postRepo{db: db}
}

func (repo *postRepo) AllPost() []model.Post {
	var posts []model.Post
	repo.db.Preload("Category").Find(&posts)
	return posts
}

func (repo *postRepo) FindByIdWithCategory(postId uint64) (model.Post, error) {
	var post model.Post
	result := repo.db.Preload("Category").First(&post, postId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return post, result.Error
	}
	return post, nil
}

func (repo *postRepo) FindById(postId uint64) (model.Post, error) {
	var post model.Post
	result := repo.db.First(&post, postId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return post, result.Error
	}
	return post, nil
}

func (repo *postRepo) Insert(post model.Post) model.Post {
	repo.db.Create(&post)
	repo.db.Preload("Category").First(&post, post.ID)
	return post
}

func (repo *postRepo) Save(post *model.Post) *gorm.DB {
	return repo.db.Save(post)
}

func (repo *postRepo) DeleteById(postId uint64) *gorm.DB {
	return repo.db.Delete(&model.Post{}, postId)
}
