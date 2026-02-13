package service

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository/mocks"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/utils"
)

// TestLogin_Success testa o login com credenciais válidas
func TestLogin_Success(t *testing.T) {
	t.Log("US0006: Login de usuário (sucesso)")
	
	// ARRANGE (Preparar): Configura o mock e dados de teste
	userRepoMock := new(mocks.UserRepositoryMock)
	os.Setenv("JWT_SECRET", "test-secret-key") // Define chave secreta para teste
	defer os.Unsetenv("JWT_SECRET")             // Limpa depois do teste
	
	service := NewAuthService(userRepoMock)
	
	// Cria uma senha hasheada para simular um usuário no banco
	hashedPassword, _ := utils.HashPassword("senha123")
	user := &domain.User{
		UserID:   1,
		UserName: "joao",
		UserType: domain.UserTypeBuyer,
		Password: hashedPassword,
	}
	
	// Configura o mock: quando FindByUsername for chamado, retorna o usuário
	userRepoMock.On("FindByUsername", "joao").Return(user, nil)
	
	// ACT (Agir): Executa o login
	response, err := service.Login("joao", "senha123")
	
	// ASSERT (Afirmar): Verifica se funcionou
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Token)       // Token foi gerado
	assert.Equal(t, uint(1), response.UserID)
	assert.Equal(t, "joao", response.UserName)
	assert.Equal(t, domain.UserTypeBuyer, response.UserType)
	
	userRepoMock.AssertExpectations(t) // Verifica que o mock foi chamado corretamente
}

// TestLogin_InvalidUsername testa login com username inexistente
func TestLogin_InvalidUsername(t *testing.T) {
	t.Log("US0006: Login de usuário (username inválido)")
	
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewAuthService(userRepoMock)
	
	// Mock retorna erro quando usuário não é encontrado
	userRepoMock.On("FindByUsername", "naoexiste").Return(nil, errors.New("not found"))
	
	response, err := service.Login("naoexiste", "senha123")
	
	// Deve retornar erro genérico por segurança
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "invalid username or password", err.Error())
}

// TestLogin_InvalidPassword testa login com senha incorreta
func TestLogin_InvalidPassword(t *testing.T) {
	t.Log("US0006: Login de usuário (senha incorreta)")
	
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewAuthService(userRepoMock)
	
	hashedPassword, _ := utils.HashPassword("senhaCorreta")
	user := &domain.User{
		UserID:   1,
		UserName: "joao",
		Password: hashedPassword,
	}
	
	userRepoMock.On("FindByUsername", "joao").Return(user, nil)
	
	// Tenta login com senha errada
	response, err := service.Login("joao", "senhaErrada")
	
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "invalid username or password", err.Error())
}

// TestValidateToken_ValidToken testa validação de token válido
func TestValidateToken_ValidToken(t *testing.T) {
	t.Log("US0006: Validação de token JWT (token válido)")
	
	userRepoMock := new(mocks.UserRepositoryMock)
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")
	
	service := NewAuthService(userRepoMock)
	authService := service.(*authServiceImpl) // Type assertion para acessar método privado
	
	// Cria um token válido manualmente
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   1,
		UserName: "joao",
		UserType: domain.UserTypeBuyer,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("test-secret-key"))
	
	// Valida o token
	validatedClaims, err := authService.ValidateToken(tokenString)
	
	assert.NoError(t, err)
	assert.NotNil(t, validatedClaims)
	assert.Equal(t, uint(1), validatedClaims.UserID)
	assert.Equal(t, "joao", validatedClaims.UserName)
	assert.Equal(t, domain.UserTypeBuyer, validatedClaims.UserType)
}

// TestValidateToken_InvalidToken testa validação com token inválido
func TestValidateToken_InvalidToken(t *testing.T) {
	t.Log("US0006: Validação de token JWT (token inválido)")
	
	userRepoMock := new(mocks.UserRepositoryMock)
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")
	
	service := NewAuthService(userRepoMock)
	authService := service.(*authServiceImpl)
	
	// Token malformado/inválido
	invalidToken := "token.invalido.aqui"
	
	claims, err := authService.ValidateToken(invalidToken)
	
	assert.Error(t, err)
	assert.Nil(t, claims)
}

// TestValidateToken_ExpiredToken testa validação com token expirado
func TestValidateToken_ExpiredToken(t *testing.T) {
	t.Log("US0006: Validação de token JWT (token expirado)")
	
	userRepoMock := new(mocks.UserRepositoryMock)
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")
	
	service := NewAuthService(userRepoMock)
	authService := service.(*authServiceImpl)
	
	// Cria token que já expirou (expirou 1 hora atrás)
	expirationTime := time.Now().Add(-1 * time.Hour)
	claims := &Claims{
		UserID:   1,
		UserName: "joao",
		UserType: domain.UserTypeBuyer,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("test-secret-key"))
	
	// Tenta validar token expirado
	validatedClaims, err := authService.ValidateToken(tokenString)
	
	assert.Error(t, err)
	assert.Nil(t, validatedClaims)
}

// TestValidateToken_WrongSecret testa validação com chave secreta diferente
func TestValidateToken_WrongSecret(t *testing.T) {
	t.Log("US0006: Validação de token JWT (chave secreta incorreta)")
	
	userRepoMock := new(mocks.UserRepositoryMock)
	
	// Cria token com uma chave
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   1,
		UserName: "joao",
		UserType: domain.UserTypeBuyer,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("secret-1"))
	
	// Tenta validar com chave diferente
	os.Setenv("JWT_SECRET", "secret-2")
	defer os.Unsetenv("JWT_SECRET")
	service2 := NewAuthService(userRepoMock)
	authService2 := service2.(*authServiceImpl)
	
	validatedClaims, err := authService2.ValidateToken(tokenString)
	
	assert.Error(t, err)
	assert.Nil(t, validatedClaims)
}

// TestNewAuthService_DefaultSecret testa criação do service sem JWT_SECRET configurado
func TestNewAuthService_DefaultSecret(t *testing.T) {
	t.Log("US0006: Criação do AuthService (sem JWT_SECRET configurado)")
	
	// Garante que não há JWT_SECRET no ambiente
	os.Unsetenv("JWT_SECRET")
	
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewAuthService(userRepoMock)
	
	// Verifica que o service foi criado (deve usar fallback)
	assert.NotNil(t, service)
	
	authService := service.(*authServiceImpl)
	assert.NotEmpty(t, authService.jwtSecret)
	// Deve usar a chave padrão de desenvolvimento
	assert.Equal(t, "my-secret-key-change-this-in-production", authService.jwtSecret)
}
