package serializer

import "github.com/abdurraufraihan/golang-blog-api/internal/model"

type CommentResponse struct {
	ID   uint   `json:"id"`
	Body string `json:"body"`
}

type CommentSerializer struct {
	Comment model.Comment
}

func (serializer *CommentSerializer) Response() CommentResponse {
	return CommentResponse{
		ID:   serializer.Comment.ID,
		Body: serializer.Comment.Body,
	}
}

type CommentsSerializer struct {
	Comments []model.Comment
}

func (serializer *CommentsSerializer) Response() []CommentResponse {
	response := []CommentResponse{}
	for _, comment := range serializer.Comments {
		commentSerializer := CommentSerializer{Comment: comment}
		response = append(response, commentSerializer.Response())
	}
	return response
}
