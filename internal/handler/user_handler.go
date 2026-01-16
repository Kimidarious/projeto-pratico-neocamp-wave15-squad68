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

func (h *UserCRUDHandler) GetAllUsers(c *gin.Context) {
    users, err := h.userService.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, users)
}


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