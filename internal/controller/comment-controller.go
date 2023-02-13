package controller

import (
	"net/http"
	"strconv"

	"github.com/abdurraufraihan/golang-blog-api/internal/dto"
	"github.com/abdurraufraihan/golang-blog-api/internal/serializer"
	"github.com/abdurraufraihan/golang-blog-api/internal/service"
	"github.com/abdurraufraihan/golang-blog-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type CommentController interface {
	Insert(context *gin.Context)
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *commentController {
	return &commentController{
		commentService: commentService,
	}
}

// InsertComment             godoc
// @Summary      Insert comment
// @Description  Responds with comment as JSON.
// @Tags         comments
// @Produce      json
// @Param        postId  path      uint  true  "Insert comment by postId"
// @Param data body dto.Comment true "Comment dto"
// @Success      201  {object}  serializer.CommentResponse
// @Router       /posts/{postId}/comments [post]
func (controller *commentController) Insert(context *gin.Context) {
	commentDto := dto.Comment{}
	err := context.ShouldBindJSON(&commentDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	postId, err := strconv.ParseUint(context.Param("postId"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "postId param not found"})
		return
	}
	tokenString := utils.GetTokenString(context)
	userId, err := utils.GetUserIDFromToken(tokenString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get userId from token"})
		return
	}
	comment := controller.commentService.Insert(commentDto, uint(postId), userId)
	serializer := serializer.CommentSerializer{Comment: comment}
	context.JSON(http.StatusCreated, serializer.Response())
}
