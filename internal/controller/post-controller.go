package controller

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/abdurraufraihan/golang-blog-api/internal/dto"
	"github.com/abdurraufraihan/golang-blog-api/internal/serializer"
	"github.com/abdurraufraihan/golang-blog-api/internal/service"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	DeleteById(context *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *postController {
	return &postController{
		postService: postService,
	}
}

// GetPosts             godoc
// @Summary      Get posts list
// @Description  Responds with the list of all posts as JSON.
// @Tags         posts
// @Produce      json
// @Success      200  {object}  serializer.PostResponse
// @Router       /posts [get]
func (controller *postController) All(context *gin.Context) {
	limit := context.Query("limit")
	offset := context.Query("offset")
	postCount, posts := controller.postService.All(limit, offset)
	serializer := serializer.PostsSerializer{Posts: posts}
	context.JSON(
		http.StatusOK, gin.H{"totalPost": postCount, "posts": serializer.Response()})
}

// GetPost             godoc
// @Summary      Get post
// @Description  Responds with post as JSON.
// @Tags         posts
// @Produce      json
// @Param        id  path      uint  true  "search post by id"
// @Success      200  {object}  serializer.PostResponse
// @Router       /posts/{id} [get]
func (controller *postController) FindById(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("postId"), 10, 64)
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

// InsertPost             godoc
// @Summary      Insert post
// @Description  Responds with post as JSON.
// @Tags         posts
// @Produce      json
// @Param data body dto.Post true "Post dto"
// @Success      201  {object}  serializer.PostResponse
// @Router       /posts [post]
func (controller *postController) Insert(context *gin.Context) {
	form := dto.Post{}
	if err := context.ShouldBind(&form); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := uploadPostImage(context, &form); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	post := controller.postService.Insert(form)
	serializer := serializer.PostSerializer{Post: post}
	context.JSON(http.StatusCreated, serializer.Response())
}

// UpdatePost             godoc
// @Summary      Update post
// @Description  Responds with post as JSON.
// @Tags         posts
// @Produce      json
// @Param        id  path      uint  true  "update post by id"
// @Param data body dto.Post true "Post dto"
// @Success      200  {object}  serializer.PostResponse
// @Router       /posts/{id} [put]
func (controller *postController) Update(context *gin.Context) {
	form := dto.Post{}
	if err := context.ShouldBind(&form); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := uploadPostImage(context, &form); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	postId, _ := strconv.ParseUint(context.Param("postId"), 10, 64)
	post, err := controller.postService.Update(postId, form)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	serializer := serializer.PostSerializer{Post: post}
	context.JSON(http.StatusOK, serializer.Response())
}

func uploadPostImage(context *gin.Context, form *dto.Post) error {
	file, _ := context.FormFile("image")
	if file != nil {
		fileName := filepath.Base(file.Filename)
		if err := context.SaveUploadedFile(file, "media/images/"+fileName); err != nil {
			return err
		}
		form.Image = "media/images/" + fileName
	}
	return nil
}

// DeletePost             godoc
// @Summary      Delete post
// @Description  Responds with post as JSON.
// @Tags         posts
// @Produce      json
// @Param        id  path      uint  true  "delete post by id"
// @Success      204
// @Router       /posts/{id} [delete]
func (controller *postController) DeleteById(context *gin.Context) {
	postId, _ := strconv.ParseUint(context.Param("postId"), 10, 64)
	result := controller.postService.DeleteById(postId)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	} else if result.RowsAffected < 1 {
		context.JSON(http.StatusNotFound, gin.H{"error": "post does not exists"})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}
