package controller

import (
	"net/http"

	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	Insert(context *gin.Context)
	All(context *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *postController {
	return &postController{
		postService: postService,
	}
}

func (controller postController) Insert(context *gin.Context) {
	postDto := dto.Post{}
	err := context.ShouldBindJSON(&postDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}
	result := controller.postService.Insert(postDto)
	context.JSON(http.StatusCreated, result)
}

func (controller postController) All(context *gin.Context) {
	posts := controller.postService.All()
	context.JSON(http.StatusOK, posts)
}
