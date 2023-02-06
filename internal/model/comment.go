package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID uint64 `json:"user"`
	User   User   `gorm:"foreignkey:UserID" json:"-"`
	PostID uint64 `json:"post"`
	Post   Post   `gorm:"foreignkey:PostID" json:"-"`
	Body   string `gorm:"type:varchar(150)" json:"body"`
}
