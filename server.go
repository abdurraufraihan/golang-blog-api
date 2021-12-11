package main

import (
	"net/http"

	"github.com/abdurraufraihan/golang-blog-api/config"
	"github.com/abdurraufraihan/golang-blog-api/controller"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.ConnectWithDB()
	postRepository repository.PostRepo       = repository.NewPostRepo(db)
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	defer config.CloseDbConnection(db)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	router.GET("/posts", postController.All)
	router.GET("/posts/:id", postController.FindById)
	router.POST("/posts", postController.Insert)
	router.Run()
}
