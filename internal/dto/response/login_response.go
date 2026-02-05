package response

import "github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"

// LoginResponse representa a resposta do endpoint de login
type LoginResponse struct {
	Token    string            `json:"token"`
	UserID   uint              `json:"user_id"`
	UserName string            `json:"user_name"`
	UserType domain.UserType   `json:"user_type"`
}
