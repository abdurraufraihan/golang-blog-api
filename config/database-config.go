package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectWithDB() *gorm.DB {
	dns := "host=localhost user=postgres dbname=goblog port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection with database")
	}
	return db
}

func CloseDbConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection with Database")
	}
	dbSql.Close()
}
