package handler

import (
	"Feature-Based/internal/post"
	"Feature-Based/internal/post/service"
	"Feature-Based/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/posts")
	{
		routes.POST("", h.CreatePost)
		routes.GET("", h.GetAllPosts)
		routes.GET(":id", h.GetPostByID)
		routes.PUT(":id", h.UpdatePost)
		routes.DELETE(":id", h.DeletePost)
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post post.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := h.service.CreatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "post created successfully"})
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (h *PostHandler) GetPostByID(c *gin.Context) {
	id, _ := utils.ParseID(c)
	post, err := h.service.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, _ := utils.ParseID(c)
	var p post.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdatePost(id, &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post updated successfully"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id, _ := utils.ParseID(c)
	if err := h.service.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
}
