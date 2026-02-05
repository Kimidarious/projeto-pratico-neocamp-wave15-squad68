package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/response"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/utils"
)

// Claims representa as informações armazenadas no token JWT
type Claims struct {
	UserID   uint            `json:"user_id"`
	UserName string          `json:"user_name"`
	UserType domain.UserType `json:"user_type"`
	jwt.RegisteredClaims
}

// AuthService interface para autenticação de usuários
type AuthService interface {
	Login(userName, password string) (*response.LoginResponse, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type authServiceImpl struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

// NewAuthService cria uma nova instância do serviço de autenticação
func NewAuthService(userRepo repository.UserRepository) AuthService {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "my-secret-key-change-this-in-production" // fallback para desenvolvimento
	}

	return &authServiceImpl{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Login autentica um usuário e retorna um token JWT
func (s *authServiceImpl) Login(userName, password string) (*response.LoginResponse, error) {
	// 1. Buscar usuário pelo username
	user, err := s.userRepo.FindByUsername(userName)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 2. Validar a senha
	if err := utils.ComparePassword(user.Password, password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 3. Gerar token JWT
	token, err := s.generateToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// 4. Retornar resposta com token e informações do usuário
	return &response.LoginResponse{
		Token:    token,
		UserID:   user.UserID,
		UserName: user.UserName,
		UserType: user.UserType,
	}, nil
}

// generateToken cria um token JWT para o usuário
func (s *authServiceImpl) generateToken(user *domain.User) (string, error) {
	// Define o tempo de expiração do token (24 horas)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Cria as claims (informações) do token
	claims := &Claims{
		UserID:   user.UserID,
		UserName: user.UserName,
		UserType: user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Cria o token com o algoritmo HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com a chave secreta
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken valida um token JWT e retorna as claims
func (s *authServiceImpl) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse e valida o token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura é o esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
