package main

import (
	"github.com/abdurraufraihan/golang-blog-api/config"
	"github.com/abdurraufraihan/golang-blog-api/route"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectWithDB()

func main() {
	defer config.CloseDbConnection(db)
	router := gin.Default()
	route.RootRoute(db, router)
	router.Run(":8000")
}
