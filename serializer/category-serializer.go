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

type CategoriesSerializer struct {
	Categories []model.Category
}

func (serializer *CategoriesSerializer) Response() []CategoryResponse {
	response := []CategoryResponse{}
	for _, category := range serializer.Categories {
		categorySerializer := CategorySerializer{Category: category}
		response = append(response, categorySerializer.Response())
	}
	return response
}
