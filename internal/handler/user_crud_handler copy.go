package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/request"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/service"
)

type UserCRUDHandler struct {
	userService service.UserService
}

func NewUserCRUDHandler(userService service.UserService) *UserCRUDHandler {
	return &UserCRUDHandler{
		userService: userService,
	}
}

// CreateUser godoc
// @Summary      Criar novo usuário
// @Description  Cria um novo usuário no sistema (comprador, vendedor ou ambos)
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      request.CreateUserRequest  true  "Dados do usuário"
// @Success      201   {object}  domain.User
// @Failure      400   {object}  map[string]string
// @Router       /users [post]
func (h *UserCRUDHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(req.UserName, req.UserType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary      Buscar usuário por ID
// @Description  Retorna os dados de um usuário específico
// @Tags         users-crud
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do usuário"
// @Success      200  {object}  domain.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserCRUDHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary      Listar todos os usuários
// @Description  Retorna a lista completa de usuários cadastrados
// @Tags         users-crud
// @Accept       json
// @Produce      json
// @Success      200  {array}   domain.User
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
func (h *UserCRUDHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary      Atualizar usuário
// @Description  Atualiza os dados de um usuário existente
// @Tags         users-crud
// @Accept       json
// @Produce      json
// @Param        id    path      int                        true  "ID do usuário"
// @Param        user  body      request.UpdateUserRequest  true  "Dados a atualizar"
// @Success      200   {object}  domain.User
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /users/{id} [put]
func (h *UserCRUDHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.UpdateUser(uint(id), req.UserName, req.UserType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Deletar usuário
// @Description  Remove um usuário do sistema
// @Tags         users-crud
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do usuário"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [delete]
func (h *UserCRUDHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}