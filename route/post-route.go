package route

import (
	"github.com/abdurraufraihan/golang-blog-api/controller"
	"github.com/abdurraufraihan/golang-blog-api/middleware"
	"github.com/abdurraufraihan/golang-blog-api/repository"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostRoute(db *gorm.DB, postRouter *gin.RouterGroup) {
	var (
		jwtService     service.JwtService        = service.NewJwtService()
		postRepository repository.PostRepo       = repository.NewPostRepo(db)
		postService    service.PostService       = service.NewPostService(postRepository)
		postController controller.PostController = controller.NewPostController(postService)
	)
	postRouter.GET("", postController.All)
	postRouter.GET("/:postId", postController.FindById)
	postRouter.POST("", middleware.AuthorizeJWT(jwtService), postController.Insert)
	postRouter.PUT("/:postId", middleware.AuthorizeJWT(jwtService), postController.Update)
}
