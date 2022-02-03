package serializer

import "github.com/abdurraufraihan/golang-blog-api/model"

type PostResponse struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Category    CategoryResponse `json:"category"`
}

type PostSerializer struct {
	Post model.Post
}

func (serializer *PostSerializer) Response() PostResponse {
	categorySerializer := CategorySerializer{Category: serializer.Post.Category}
	return PostResponse{
		ID:          serializer.Post.ID,
		Title:       serializer.Post.Title,
		Description: serializer.Post.Description,
		Category:    categorySerializer.Response(),
	}
}

type PostsSerializer struct {
	Posts []model.Post
}

func (serializer *PostsSerializer) Response() []PostResponse {
	response := []PostResponse{}
	for _, post := range serializer.Posts {
		serializer := PostSerializer{Post: post}
		response = append(response, serializer.Response())
	}
	return response
}
