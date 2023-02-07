package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body   string `gorm:"type:varchar(150)" json:"body"`
	UserID uint
	PostID uint
}
