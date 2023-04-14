package route

import (
	"github.com/abdurraufraihan/golang-blog-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootRoute(db *gorm.DB, router *gin.Engine, logger *logger.Logger) {
	router.Static("/media", "/media")
	apiRouter := router.Group("/api/v1")
	postRouter := apiRouter.Group("/posts")
	PostRoute(db, postRouter)
	commentRouter := apiRouter.Group("/posts/:postId/comments")
	CommentRoute(db, commentRouter, logger)
	categoryRouter := apiRouter.Group("/categories")
	CategoryRoute(db, categoryRouter, logger)
	authRouter := apiRouter.Group("/auth")
	AuthRoute(db, authRouter, logger)
}
