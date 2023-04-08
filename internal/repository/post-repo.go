package repository

import (
	"errors"
	"strconv"

	"github.com/abdurraufraihan/golang-blog-api/internal/model"
	"gorm.io/gorm"
)

type PostRepo interface {
	PostCount() int64
	AllPost(limit string, offset string) []model.Post
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

func (repo *postRepo) PostCount() int64 {
	var count int64
	repo.db.Model(&model.Post{}).Count(&count)
	return count
}

func (repo *postRepo) AllPost(limit string, offset string) []model.Post {
	postLimit, err := strconv.Atoi(limit)
	if err != nil || postLimit == 0 {
		postLimit = -1
	}
	postOffset, err := strconv.Atoi(offset)
	if err != nil || postOffset == 0 {
		postOffset = -1
	}
	var posts []model.Post
	repo.db.Limit(postLimit).Offset(postOffset).
		Preload("Category").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Limit(5).Order("id desc")
	}).Find(&posts)
	return posts
}

func (repo *postRepo) FindByIdWithCategory(postId uint64) (model.Post, error) {
	var post model.Post
	result := repo.db.Preload("Category").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Limit(5).Order("id desc")
		}).First(&post, postId)
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
