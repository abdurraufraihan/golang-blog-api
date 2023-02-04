package main

import (
	"github.com/abdurraufraihan/golang-blog-api/adapter"
	"github.com/abdurraufraihan/golang-blog-api/internal/route"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = adapter.ConnectWithDB()

func main() {
	defer adapter.CloseDbConnection(db)
	router := gin.Default()
	route.RootRoute(db, router)
	router.Run(":8000")
}
