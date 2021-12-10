package dto

type Post struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  uint   `json:"category" binding:"required"`
}
