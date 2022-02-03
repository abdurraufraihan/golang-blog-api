package serializer

import "github.com/abdurraufraihan/golang-blog-api/model"

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategorySerializer struct {
	Category model.Category
}

func (serializer *CategorySerializer) Response() CategoryResponse {
	return CategoryResponse{
		ID:   serializer.Category.ID,
		Name: serializer.Category.Name,
	}
}
