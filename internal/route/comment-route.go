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

func CommentRoute(
	db *gorm.DB, CommentRouter *gin.RouterGroup, logger *logger.Logger,
) {
	var (
		commentRepository repository.CommentRepo = repository.NewCommentRepo(db)
		commentService    service.CommentService = service.
					NewCommentService(commentRepository)
		commentController controller.CommentController = controller.
					NewCommentController(commentService, logger)
	)
	CommentRouter.GET("", commentController.All)
	CommentRouter.POST("", middleware.AuthorizeJWT(), commentController.Insert)
}
