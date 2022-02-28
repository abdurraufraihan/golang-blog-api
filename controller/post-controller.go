package controller

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/abdurraufraihan/golang-blog-api/dto"
	"github.com/abdurraufraihan/golang-blog-api/serializer"
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
	form := dto.Post{}
	if err := context.ShouldBind(&form); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if form.Image != "" {
		file, _ := context.FormFile("image")
		fileName := filepath.Base(file.Filename)
		if err := context.SaveUploadedFile(file, "media/images/"+fileName); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		form.Image = "media/images/" + fileName
	}
	post := controller.postService.Insert(form)
	serializer := serializer.PostSerializer{Post: post}
	context.JSON(http.StatusCreated, serializer.Response())
}

func (controller *postController) All(context *gin.Context) {
	posts := controller.postService.All()
	serializer := serializer.PostsSerializer{Posts: posts}
	context.JSON(http.StatusOK, serializer.Response())
}

func (controller *postController) FindById(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No param id was found"})
		return
	}
	post, err := controller.postService.FindById(postId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	serializer := serializer.PostSerializer{Post: post}
	context.JSON(http.StatusOK, serializer.Response())
}
