package post

import (
	"ddd/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase *UseCase
}

func NewPostHandler(useCase *UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func RegisterPostRoutes(router *gin.Engine, handler *Handler) {
	routes := router.Group("/posts")
	{
		routes.POST("", handler.CreatePost)
		routes.GET("", handler.GetAllPosts)
		routes.GET(":id", handler.GetPostByID)
		routes.PUT(":id", handler.UpdatePost)
		routes.DELETE(":id", handler.DeletePost)
	}
}
func (h *Handler) CreatePost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.useCase.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "post created successfully",
	})
}

func (h *Handler) GetAllPosts(c *gin.Context) {
	posts, err := h.useCase.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPostByID(c *gin.Context) {
	id, _ := utils.ParseID(c)
	post, err := h.useCase.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) UpdatePost(c *gin.Context) {
	id, _ := utils.ParseID(c)
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.useCase.UpdatePost(id, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "post updated successfully",
	})
}

func (h *Handler) DeletePost(c *gin.Context) {
	id, _ := utils.ParseID(c)
	if err := h.useCase.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
