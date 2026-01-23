package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository/mocks"
)

func TestCreateUser_Success(t *testing.T) {
	t.Log("US0013: Cadastro de novo usuário (sucesso)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userName := "joao"
	userType := domain.UserTypeBuyer

	userRepoMock.On("FindByUsername", userName).Return(nil, errors.New("not found"))
	userRepoMock.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)

	user, err := service.CreateUser(userName, userType)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userName, user.UserName)
	assert.Equal(t, domain.UserTypeBuyer, user.UserType)

	userRepoMock.AssertExpectations(t)
}

func TestCreateUser_DuplicateUsername(t *testing.T) {
	t.Log("US0013: Cadastro de novo usuário (username duplicado)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userName := "joao"
	userType := domain.UserTypeBuyer

	existingUser := &domain.User{
		UserID:   1,
		UserName: userName,
		UserType: domain.UserTypeBuyer,
	}
	userRepoMock.On("FindByUsername", userName).Return(existingUser, nil)

	user, err := service.CreateUser(userName, userType)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "username already exists", err.Error())

	userRepoMock.AssertNotCalled(t, "Create")
}

func TestCreateUser_InvalidUserType(t *testing.T) {
	t.Log("US0013: Cadastro de novo usuário (tipo inválido)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userName := "joao"
	userType := domain.UserType("INVALID")

	userRepoMock.On("FindByUsername", userName).Return(nil, errors.New("not found"))

	user, err := service.CreateUser(userName, userType)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "invalid user type")
}

func TestDeleteUser_Success(t *testing.T) {
	t.Log("US0013: Excluir usuário (sucesso)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userID := uint(1)

	userRepoMock.On("ExistsByID", userID).Return(true)
	userRepoMock.On("Delete", userID).Return(nil)

	err := service.DeleteUser(userID)

	assert.NoError(t, err)
	userRepoMock.AssertExpectations(t)
}

func TestDeleteUser_NotFound(t *testing.T) {
	t.Log("US0013: Excluir usuário (não encontrado)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userID := uint(999)

	userRepoMock.On("ExistsByID", userID).Return(false)

	err := service.DeleteUser(userID)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	userRepoMock.AssertNotCalled(t, "Delete")
}

func TestGetUserByID_Success(t *testing.T) {
	t.Log("US0013: Buscar usuário por ID (sucesso)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userID := uint(1)
	expected := &domain.User{
		UserID:   userID,
		UserName: "maria",
		UserType: domain.UserTypeBuyer,
	}

	userRepoMock.On("FindByID", userID).Return(expected, nil)

	user, err := service.GetUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, expected, user)
	userRepoMock.AssertExpectations(t)
}

func TestGetAllUsers_Success(t *testing.T) {
	t.Log("US0013: Listar todos os usuários")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	users := []*domain.User{
		{UserID: 1, UserName: "ana", UserType: domain.UserTypeBuyer},
		{UserID: 2, UserName: "joao", UserType: domain.UserTypeSeller},
	}

	userRepoMock.On("FindAll").Return(users, nil)

	result, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, users, result)
	userRepoMock.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	t.Log("US0013: Atualizar usuário (sucesso)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userID := uint(1)
	existing := &domain.User{
		UserID:   userID,
		UserName: "joao",
		UserType: domain.UserTypeBuyer,
	}

	userRepoMock.On("FindByID", userID).Return(existing, nil)
	userRepoMock.On("Update", mock.AnythingOfType("*domain.User")).Return(nil)

	user, err := service.UpdateUser(userID, "carlos", domain.UserTypeSeller)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "carlos", user.UserName)
	assert.Equal(t, domain.UserTypeSeller, user.UserType)
	userRepoMock.AssertExpectations(t)
}

func TestUpdateUser_InvalidUserType(t *testing.T) {
	t.Log("US0013: Atualizar usuário (tipo inválido)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userID := uint(1)
	existing := &domain.User{
		UserID:   userID,
		UserName: "joao",
		UserType: domain.UserTypeBuyer,
	}

	userRepoMock.On("FindByID", userID).Return(existing, nil)

	user, err := service.UpdateUser(userID, "carlos", domain.UserType("INVALID"))

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "invalid user type")
	userRepoMock.AssertNotCalled(t, "Update")
}

func TestUpdateUser_NotFound(t *testing.T) {
	t.Log("US0013: Atualizar usuário (não encontrado)")
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)

	userID := uint(999)

	userRepoMock.On("FindByID", userID).Return(nil, errors.New("user not found"))

	user, err := service.UpdateUser(userID, "carlos", domain.UserTypeSeller)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "user not found", err.Error())
	userRepoMock.AssertNotCalled(t, "Update")
}
