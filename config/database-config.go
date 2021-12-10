package config

import (
	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectWithDB() *gorm.DB {
	dns := "host=localhost user=postgres password=123456 dbname=goblog port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection with database")
	}
	db.AutoMigrate(&model.Post{}, &model.Category{})
	return db
}

func CloseDbConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection with Database")
	}
	dbSql.Close()
}
