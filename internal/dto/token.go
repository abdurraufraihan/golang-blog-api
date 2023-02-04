package dto

type Token struct {
	Token string `json:"token" binding:"required"`
}
