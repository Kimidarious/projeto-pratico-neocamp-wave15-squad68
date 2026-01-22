package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository/mocks"
)

func TestFollowUser_FollowedUserNotFound(t *testing.T) {
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	
	service := NewFollowService(followRepoMock, userRepoMock)
	
	followerID := uint(1)
	followedID := uint(999) // usuário inexistente
	
	userRepoMock.On("ExistsByID", followerID).Return(true)
	
	userRepoMock.On("ExistsByID", followedID).Return(false)
	
	err := service.FollowUser(followerID, followedID)
	
	assert.Error(t, err)
	assert.Equal(t, "followed user not found", err.Error())
	
	userRepoMock.AssertExpectations(t)
	followRepoMock.AssertNotCalled(t, "Create") // Não deve ter chamado Create
}

func TestFollowUser_Success(t *testing.T) {
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	
	service := NewFollowService(followRepoMock, userRepoMock)
	
	followerID := uint(1)
	followedID := uint(2)
	
	userRepoMock.On("ExistsByID", followerID).Return(true)
	userRepoMock.On("ExistsByID", followedID).Return(true)
	
	followRepoMock.On("IsFollowing", followerID, followedID).Return(false)
	
	followRepoMock.On("Create", mock.AnythingOfType("*domain.Follow")).Return(nil)
	
	err := service.FollowUser(followerID, followedID)
	
	assert.NoError(t, err)
	
	userRepoMock.AssertExpectations(t)
	followRepoMock.AssertExpectations(t)
	followRepoMock.AssertCalled(t, "Create", mock.AnythingOfType("*domain.Follow"))
}

func TestFollowUser_CannotFollowYourself(t *testing.T) {
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	
	service := NewFollowService(followRepoMock, userRepoMock)
	
	userID := uint(1)
	
	err := service.FollowUser(userID, userID)
	
	assert.Error(t, err)
	assert.Equal(t, "cannot follow yourself", err.Error())
	
	userRepoMock.AssertNotCalled(t, "ExistsByID")
	followRepoMock.AssertNotCalled(t, "Create")
}

func TestFollowUser_AlreadyFollowing(t *testing.T) {
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	
	service := NewFollowService(followRepoMock, userRepoMock)
	
	followerID := uint(1)
	followedID := uint(2)
	
	userRepoMock.On("ExistsByID", followerID).Return(true)
	userRepoMock.On("ExistsByID", followedID).Return(true)
	
	followRepoMock.On("IsFollowing", followerID, followedID).Return(true)
	
	err := service.FollowUser(followerID, followedID)
	
	assert.Error(t, err)
	assert.Equal(t, "already following this user", err.Error())
	
	followRepoMock.AssertNotCalled(t, "Create") // Não deve criar
}

func TestGetFollowersList_OrderNameAsc(t *testing.T) {
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	
	service := NewFollowService(followRepoMock, userRepoMock)
	
	userID := uint(1)
	order := "name_asc"
	
	user := &domain.User{
		UserID:   userID,
		UserName: "maria",
		UserType: domain.UserTypeSeller,
	}
	
	followers := []*domain.User{
		{UserID: 2, UserName: "ana"},
		{UserID: 3, UserName: "carlos"},
		{UserID: 4, UserName: "joao"},
	}
	
	userRepoMock.On("FindByID", userID).Return(user, nil)
	followRepoMock.On("GetFollowers", userID, order).Return(followers, nil)
	
	result, err := service.GetFollowersList(userID, order)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, len(result.Followers))
	
	assert.Equal(t, "ana", result.Followers[0].UserName)
	assert.Equal(t, "carlos", result.Followers[1].UserName)
	assert.Equal(t, "joao", result.Followers[2].UserName)
	
	userRepoMock.AssertExpectations(t)
	followRepoMock.AssertExpectations(t)
}

func TestGetFollowersList_OrderNameDesc(t *testing.T) {
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	
	service := NewFollowService(followRepoMock, userRepoMock)
	
	userID := uint(1)
	order := "name_desc"
	
	user := &domain.User{
		UserID:   userID,
		UserName: "maria",
		UserType: domain.UserTypeSeller,
	}
	
	followers := []*domain.User{
		{UserID: 4, UserName: "joao"},
		{UserID: 3, UserName: "carlos"},
		{UserID: 2, UserName: "ana"},
	}
	
	userRepoMock.On("FindByID", userID).Return(user, nil)
	followRepoMock.On("GetFollowers", userID, order).Return(followers, nil)
	
	result, err := service.GetFollowersList(userID, order)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, len(result.Followers))
	
	assert.Equal(t, "joao", result.Followers[0].UserName)
	assert.Equal(t, "carlos", result.Followers[1].UserName)
	assert.Equal(t, "ana", result.Followers[2].UserName)
	
	userRepoMock.AssertExpectations(t)
	followRepoMock.AssertExpectations(t)
}

func TestUnfollowUser_NotFollowing(t *testing.T) {
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	
	service := NewFollowService(followRepoMock, userRepoMock)
	
	followerID := uint(1)
	followedID := uint(2)
	
	followRepoMock.On("IsFollowing", followerID, followedID).Return(false)
	
	err := service.UnfollowUser(followerID, followedID)
	
	assert.Error(t, err)
	assert.Equal(t, "not following this user", err.Error())
	
	followRepoMock.AssertNotCalled(t, "Delete")
}