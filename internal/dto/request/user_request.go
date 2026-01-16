package request

import "github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"

type CreateUserRequest struct {
    UserName string            `json:"user_name" binding:"required,min=3,max=15"`
    UserType domain.UserType   `json:"user_type" binding:"required"`
}

type UpdateUserRequest struct {
    UserName string            `json:"user_name,omitempty" binding:"omitempty,min=3,max=15"`
    UserType domain.UserType   `json:"user_type,omitempty"`
}