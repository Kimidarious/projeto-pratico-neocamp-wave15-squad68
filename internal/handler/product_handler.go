package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/request"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/service"
)

type ProductHandler struct {
	postService service.PostService
}

func NewProductHandler(postService service.PostService) *ProductHandler {
	return &ProductHandler{
		postService: postService,
	}
}


func (h *ProductHandler) CreatePost(c *gin.Context) {
	var req request.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postService.CreatePost(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}


func (h *ProductHandler) CreatePromoPost(c *gin.Context) {
	var req request.CreatePromoPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postService.CreatePromoPost(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}


func (h *ProductHandler) GetFollowedPosts(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	order := c.DefaultQuery("order", "date_desc")

	result, err := h.postService.GetFollowedPosts(uint(userID), order)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}


func (h *ProductHandler) CountPromoProducts(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	result, err := h.postService.CountPromoProducts(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}