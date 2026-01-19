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

// US-0005: Create Post
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

// US-0010: Create Promo Post
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

// US-0006 + US-0009: Get Followed Posts
func (h *ProductHandler) GetFollowedPosts(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

// US-0011: Count Promo Products
func (h *ProductHandler) CountPromoProducts(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

// US-0012: Listar produtos promocionais de um vendedor
// @Summary      Listar produtos promocionais
// @Description  Retorna todos os produtos em promoção de um vendedor
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  response.PostListResponse
// @Failure      404  {object}  map[string]string
// @Router       /products/{id}/promos [get]
func (h *ProductHandler) GetPromoPostsByUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	result, err := h.postService.GetPromoPostsByUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}