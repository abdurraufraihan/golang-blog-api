package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootRoute(db *gorm.DB, router *gin.Engine) {
	router.Static("/media", "/media")
	apiRouter := router.Group("/api/v1")
	postRouter := apiRouter.Group("/posts")
	PostRoute(db, postRouter)
	commentRouter := apiRouter.Group("/posts/:postId/comments")
	CommentRoute(db, commentRouter)
	categoryRouter := apiRouter.Group("/categories")
	CategoryRoute(db, categoryRouter)
	authRouter := apiRouter.Group("/auth")
	AuthRoute(db, authRouter)
}
