package request

// LoginRequest representa os dados necessários para fazer login
type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
