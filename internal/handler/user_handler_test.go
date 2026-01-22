package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/service"
)


type UserHandler struct {
	followService service.FollowService
}

func NewUserHandler(followService service.FollowService) *UserHandler {
	return &UserHandler{
		followService: followService,
	}
}

// FollowUser godoc
// @Summary      Seguir um vendedor
// @Description  Permite que um usuário siga um vendedor (US-0001)
// @Tags         users-follow
// @Accept       json
// @Produce      json
// @Param        id              path      int  true  "ID do usuário que vai seguir"
// @Param        userIdToFollow  path      int  true  "ID do usuário a ser seguido"
// @Success      200             {object}  map[string]string
// @Failure      400             {object}  map[string]string
// @Router       /users/{id}/follow/{userIdToFollow} [post]
func (h *UserHandler) FollowUser(c *gin.Context) {
	followerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followedID, err := strconv.ParseUint(c.Param("userIdToFollow"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to follow"})
		return
	}

	if err := h.followService.FollowUser(uint(followerID), uint(followedID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully followed user"})
}

// UnfollowUser godoc
// @Summary      Deixar de seguir um vendedor
// @Description  Remove o vínculo de follow entre usuários (US-0007)
// @Tags         users-follow
// @Accept       json
// @Produce      json
// @Param        id              path      int  true  "ID do usuário que vai deixar de seguir"
// @Param        userIdToFollow  path      int  true  "ID do usuário a ser deixado de seguir"
// @Success      200             {object}  map[string]string
// @Failure      400             {object}  map[string]string
// @Router       /users/{id}/unfollow/{userIdToFollow} [post]
func (h *UserHandler) UnfollowUser(c *gin.Context) {
	followerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followedID, err := strconv.ParseUint(c.Param("userIdToFollow"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to unfollow"})
		return
	}

	if err := h.followService.UnfollowUser(uint(followerID), uint(followedID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully unfollowed user"})
}

// GetFollowersCount godoc
// @Summary      Obter contagem de seguidores
// @Description  Retorna o número de seguidores de um vendedor (US-0002)
// @Tags         users-followers
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do usuário"
// @Success      200  {object}  response.FollowersCountResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id}/followers/count [get]
func (h *UserHandler) GetFollowersCount(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	result, err := h.followService.GetFollowersCount(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}


// GetFollowersList godoc
// @Summary      Listar seguidores
// @Description  Retorna a lista de usuários que seguem um vendedor (US-0003, US-0008)
// @Tags         users-followers
// @Accept       json
// @Produce      json
// @Param        id     path      int     true   "ID do usuário"
// @Param        order  query     string  false  "Ordenação: name_asc ou name_desc"  default(name_asc)
// @Success      200    {object}  response.FollowersListResponse
// @Failure      400    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /users/{id}/followers/list [get]
func (h *UserHandler) GetFollowersList(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	order := c.DefaultQuery("order", "name_asc")

	result, err := h.followService.GetFollowersList(uint(userID), order)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetFollowedList godoc
// @Summary      Listar seguidos
// @Description  Retorna a lista de vendedores que um usuário segue (US-0004, US-0008)
// @Tags         Follow
// @Accept       json
// @Produce      json
// @Param        id     path      int     true   "ID do usuário"
// @Param        order  query     string  false  "Ordenação: name_asc ou name_desc"  default(name_asc)
// @Success      200    {object}  response.FollowedListResponse
// @Failure      400    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /users/{id}/followed/list [get]
func (h *UserHandler) GetFollowedList(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	order := c.DefaultQuery("order", "name_asc")

	result, err := h.followService.GetFollowedList(uint(userID), order)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}