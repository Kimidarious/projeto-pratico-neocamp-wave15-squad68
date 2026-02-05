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

// CreatePost godoc
// @Summary      Criar publicação
// @Description  Cria uma nova publicação de produto (US-0005)
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        post  body  request.CreatePostRequest  true  "Dados da publicação"
// @Success      201
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Security     BearerAuth
// @Router       /products/post [post]
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

// CreatePromoPost godoc
// @Summary      Criar publicação promocional
// @Description  Cria uma publicação de produto com promoção (US-0010)
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        post  body  request.CreatePromoPostRequest  true  "Dados da publicação promocional"
// @Success      201
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Security     BearerAuth
// @Router       /products/promo-post [post]
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

// GetFollowedPosts godoc
// @Summary      Obter feed de publicações
// @Description  Retorna publicações dos vendedores seguidos nas últimas 2 semanas (US-0006, US-0009)
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id     path      int     true   "ID do usuário"
// @Param        order  query     string  false  "Ordenação: date_asc ou date_desc"  default(date_desc)
// @Success      200    {object}  response.PostListResponse
// @Failure      400    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /products/followed/{id}/list [get]
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

// CountPromoProducts godoc
// @Summary      Contar produtos em promoção
// @Description  Retorna a quantidade de produtos promocionais de um vendedor (US-0011)
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do usuário"
// @Success      200  {object}  response.PromoCountResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /products/{id}/countPromo [get]
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

// GetPromoPostsByUser godoc
// @Summary      Listar produtos promocionais
// @Description  Retorna todos os produtos em promoção de um vendedor (US-0012)
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do usuário"
// @Success      200  {object}  response.PostListResponse
// @Failure      400  {object}  map[string]string
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