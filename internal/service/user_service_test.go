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
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)
	
	userName := "joao"
	userType := "BUYER"
	
	userRepoMock.On("ExistsByUsername", userName).Return(false)
	
	userRepoMock.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)
	
	user, err := service.CreateUser(userName, userType)
	
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userName, user.UserName)
	assert.Equal(t, domain.UserType(userType), user.UserType)
	
	userRepoMock.AssertExpectations(t)
}

func TestCreateUser_DuplicateUsername(t *testing.T) {
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)
	
	userName := "joao"
	userType := "BUYER"
	
	userRepoMock.On("ExistsByUsername", userName).Return(true)
	
	user, err := service.CreateUser(userName, userType)
	
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "username already exists", err.Error())
	
	userRepoMock.AssertNotCalled(t, "Create")
}

func TestCreateUser_InvalidUserType(t *testing.T) {
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)
	
	userName := "joao"
	userType := "INVALID"
	
	userRepoMock.On("ExistsByUsername", userName).Return(false)
	
	user, err := service.CreateUser(userName, userType)
	
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "invalid user type")
}

func TestDeleteUser_Success(t *testing.T) {
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)
	
	userID := uint(1)
	
	userRepoMock.On("Delete", userID).Return(nil)
	
	err := service.DeleteUser(userID)
	
	assert.NoError(t, err)
	userRepoMock.AssertExpectations(t)
}

func TestDeleteUser_NotFound(t *testing.T) {
	userRepoMock := new(mocks.UserRepositoryMock)
	service := NewUserService(userRepoMock)
	
	userID := uint(999)
	
	userRepoMock.On("Delete", userID).Return(errors.New("user not found"))
	
	err := service.DeleteUser(userID)
	
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}