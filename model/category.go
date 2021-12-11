package model

type Category struct {
	ID    uint   `gorm:"primary_key:auto_increment" json:"id"`
	Name  string `gorm:"type:varchar(100)" json:"name"`
	Posts []Post
}
