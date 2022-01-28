package controller

import (
	"net/http"
	"strconv"

	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/model"
	"github.com/abdurraufraihan/golang-blog-api/service"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	Insert(context *gin.Context)
	All(context *gin.Context)
	FindById(context *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *postController {
	return &postController{
		postService: postService,
	}
}

func (controller *postController) Insert(context *gin.Context) {
	postDto := dto.Post{}
	err := context.ShouldBindJSON(&postDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}
	result := controller.postService.Insert(postDto)
	context.JSON(http.StatusCreated, result)
}

func (controller *postController) All(context *gin.Context) {
	posts := controller.postService.All()
	context.JSON(http.StatusOK, posts)
}

func (controller *postController) FindById(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No param id was found"})
	}
	var post model.Post = controller.postService.FindById(postId)
	if (post == model.Post{}) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
	} else {
		context.JSON(http.StatusOK, post)
	}
}
