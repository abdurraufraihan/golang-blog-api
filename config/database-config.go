package config

import (
	"fmt"

	"github.com/abdurraufraihan/golang-blog-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     int
	sslmode  string
	timezone string
}

func getDBConfig() *dbConfig {
	return &dbConfig{
		host:     "localhost",
		user:     "postgres",
		password: "123456",
		dbname:   "goblog",
		port:     5432,
		sslmode:  "disable",
		timezone: "UTC",
	}
}

func ConnectWithDB() *gorm.DB {
	dbConfig := getDBConfig()
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.host,
		dbConfig.user,
		dbConfig.password,
		dbConfig.dbname,
		dbConfig.port,
		dbConfig.sslmode,
		dbConfig.timezone,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection with database")
	}
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Post{})
	return db
}

func CloseDbConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection with Database")
	}
	dbSql.Close()
}
