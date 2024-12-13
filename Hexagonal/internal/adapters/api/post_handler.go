package api

import (
	"Hexagonal/internal/core/post"
	"Hexagonal/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostHandler struct {
	useCase *post.PostUseCase
}

func NewPostHandler(useCase *post.PostUseCase) *PostHandler {
	return &PostHandler{useCase: useCase}
}

func (h *PostHandler) RegisterRoutes(router *gin.Engine) {
	r := router.Group("/posts")
	{
		r.POST("", h.CreatePost)
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var p domain.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	if err := h.useCase.StorePost(ctx, &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "post created successfully"})
}
