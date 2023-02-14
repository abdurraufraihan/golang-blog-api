package serializer

import (
	"time"

	"github.com/abdurraufraihan/golang-blog-api/internal/model"
)

type PostResponse struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Category    CategoryResponse  `json:"category"`
	Comments    []CommentResponse `json:"comments"`
}

type PostSerializer struct {
	Post model.Post
}

func (serializer *PostSerializer) Response() PostResponse {
	categorySerializer := CategorySerializer{Category: serializer.Post.Category}
	commentSerializer := CommentsSerializer{Comments: serializer.Post.Comments}
	return PostResponse{
		ID:          serializer.Post.ID,
		Title:       serializer.Post.Title,
		Description: serializer.Post.Description,
		Image:       serializer.Post.Image,
		CreatedAt:   serializer.Post.CreatedAt,
		UpdatedAt:   serializer.Post.UpdatedAt,
		Category:    categorySerializer.Response(),
		Comments:    commentSerializer.Response(),
	}
}

type PostsSerializer struct {
	Posts []model.Post
}

func (serializer *PostsSerializer) Response() []PostResponse {
	response := []PostResponse{}
	for _, post := range serializer.Posts {
		postSerializer := PostSerializer{Post: post}
		response = append(response, postSerializer.Response())
	}
	return response
}
