package dto

type CommentDto struct {
	Body string `json:"body" binding:"required,max=100"`
}
