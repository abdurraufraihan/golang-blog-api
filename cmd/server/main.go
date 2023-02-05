package main

import (
	"os"

	"github.com/abdurraufraihan/golang-blog-api/adapter"
	"github.com/abdurraufraihan/golang-blog-api/docs"
	"github.com/abdurraufraihan/golang-blog-api/internal/route"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var db *gorm.DB = adapter.ConnectWithDB()

// @title           Gin Blog Api
// @version         1.0
// @description     A Blog API in Go using Gin framework.

// @contact.name   Abdur Rauf Raihan
// @contact.url    https://linkedin.com/abdurraufraihan
// @contact.email  abdurraufraihan@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

func main() {
	defer adapter.CloseDbConnection(db)
	router := gin.Default()
	route.RootRoute(db, router)
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + os.Getenv("APP_PORT"))
}
