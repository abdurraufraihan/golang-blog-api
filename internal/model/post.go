package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(100)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`
	CategoryID  uint      `json:"category"`
	Category    Category  `gorm:"foreignkey:CategoryID" json:"-"`
	Comments    []Comment `json:"comments"`
}
