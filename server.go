package main

import (
	"net/http"

	"github.com/abdurraufraihan/golang-blog-api/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectWithDB()

func main() {
	defer config.CloseDbConnection(db)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	router.Run()
}
