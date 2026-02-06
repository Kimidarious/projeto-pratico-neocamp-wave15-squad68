package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/request"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/response"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/service"
)

// MockAuthService é um mock do AuthService
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(userName, password string) (*response.LoginResponse, error) {
	args := m.Called(userName, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.LoginResponse), args.Error(1)
}

func (m *MockAuthService) ValidateToken(tokenString string) (*service.Claims, error) {
	args := m.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.Claims), args.Error(1)
}

// TestLogin_Success testa o endpoint de login com credenciais válidas
func TestLogin_Success(t *testing.T) {
	// ARRANGE
	gin.SetMode(gin.TestMode) // Modo de teste do Gin
	
	mockService := new(MockAuthService)
	handler := NewAuthHandler(mockService)
	
	// Configura resposta esperada do mock
	expectedResponse := &response.LoginResponse{
		Token:    "token-jwt-aqui",
		UserID:   1,
		UserName: "joao",
		UserType: domain.UserTypeBuyer,
	}
	mockService.On("Login", "joao", "senha123").Return(expectedResponse, nil)
	
	// Cria requisição HTTP simulada
	loginReq := request.LoginRequest{
		UserName: "joao",
		Password: "senha123",
	}
	body, _ := json.Marshal(loginReq)
	
	// Cria recorder para capturar a resposta
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")
	
	// ACT
	handler.Login(c)
	
	// ASSERT
	assert.Equal(t, http.StatusOK, w.Code) // Status 200
	
	var actualResponse response.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Token, actualResponse.Token)
	assert.Equal(t, expectedResponse.UserID, actualResponse.UserID)
	assert.Equal(t, expectedResponse.UserName, actualResponse.UserName)
	
	mockService.AssertExpectations(t)
}

// TestLogin_InvalidJSON testa com JSON malformado
func TestLogin_InvalidJSON(t *testing.T) {
	// ARRANGE
	gin.SetMode(gin.TestMode)
	
	mockService := new(MockAuthService)
	handler := NewAuthHandler(mockService)
	
	// JSON inválido
	invalidJSON := []byte(`{"username": "joao", "password":}`)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(invalidJSON))
	c.Request.Header.Set("Content-Type", "application/json")
	
	// ACT
	handler.Login(c)
	
	// ASSERT
	assert.Equal(t, http.StatusBadRequest, w.Code) // Status 400
	
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["error"], "invalid")
}

// TestLogin_InvalidCredentials testa com credenciais inválidas
func TestLogin_InvalidCredentials(t *testing.T) {
	// ARRANGE
	gin.SetMode(gin.TestMode)
	
	mockService := new(MockAuthService)
	handler := NewAuthHandler(mockService)
	
	// Mock retorna erro
	mockService.On("Login", "joao", "senhaErrada").
		Return(nil, errors.New("invalid username or password"))
	
	loginReq := request.LoginRequest{
		UserName: "joao",
		Password: "senhaErrada",
	}
	body, _ := json.Marshal(loginReq)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")
	
	// ACT
	handler.Login(c)
	
	// ASSERT
	assert.Equal(t, http.StatusUnauthorized, w.Code) // Status 401
	
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "invalid username or password", response["error"])
	
	mockService.AssertExpectations(t)
}
