package route

import (
	"github.com/abdurraufraihan/golang-blog-api/internal/controller"
	"github.com/abdurraufraihan/golang-blog-api/internal/middleware"
	"github.com/abdurraufraihan/golang-blog-api/internal/repository"
	"github.com/abdurraufraihan/golang-blog-api/internal/service"
	"github.com/abdurraufraihan/golang-blog-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostRoute(db *gorm.DB, postRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		postRepository repository.PostRepo = repository.NewPostRepo(db)
		postService    service.PostService = service.
				NewPostService(postRepository)
		postController controller.PostController = controller.
				NewPostController(postService, logger)
	)
	postRouter.GET("", postController.All)
	postRouter.GET("/:postId", postController.FindById)
	postRouter.POST("", middleware.AuthorizeJWT(), postController.Insert)
	postRouter.PUT("/:postId", middleware.AuthorizeJWT(), postController.Update)
	postRouter.DELETE(
		"/:postId", middleware.AuthorizeJWT(), postController.DeleteById)
}
