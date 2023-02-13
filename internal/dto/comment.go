package dto

type Comment struct {
	Body string `json:"body" binding:"required,max=100"`
}
