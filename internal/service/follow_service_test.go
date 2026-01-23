package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository/mocks"
)

func TestFollowUser_FollowedUserNotFound(t *testing.T) {
	t.Log("US0001: Follow - erro quando vendedor não existe")
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)

	service := NewFollowService(followRepoMock, userRepoMock)

	followerID := uint(1)
	followedID := uint(999)

	userRepoMock.On("ExistsByID", followerID).Return(true)
	userRepoMock.On("ExistsByID", followedID).Return(false)

	err := service.FollowUser(followerID, followedID)

	assert.Error(t, err)
	assert.Equal(t, "followed user not found", err.Error())

	userRepoMock.AssertExpectations(t)
	followRepoMock.AssertNotCalled(t, "Create")
}

func TestFollowUser_Success(t *testing.T) {
	t.Log("US0001: Follow - sucesso ao seguir vendedor")
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
}

func TestFollowUser_CannotFollowYourself(t *testing.T) {
	t.Log("US0001: Follow - bloqueia seguir a si mesmo")
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
	t.Log("US0001: Follow - erro quando já segue o vendedor")
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
	followRepoMock.AssertNotCalled(t, "Create")
}

func TestGetFollowersList_OrderNameAsc(t *testing.T) {
	t.Log("US0003/US0008: Lista seguidores ordenada por nome asc")
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
		{UserID: 2, UserName: "ana", UserType: domain.UserTypeBuyer},
		{UserID: 3, UserName: "carlos", UserType: domain.UserTypeBuyer},
		{UserID: 4, UserName: "joao", UserType: domain.UserTypeBuyer},
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
	t.Log("US0003/US0008: Lista seguidores ordenada por nome desc")
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
		{UserID: 4, UserName: "joao", UserType: domain.UserTypeBuyer},
		{UserID: 3, UserName: "carlos", UserType: domain.UserTypeBuyer},
		{UserID: 2, UserName: "ana", UserType: domain.UserTypeBuyer},
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
	t.Log("US0007: Unfollow - erro quando não segue")
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

func TestUnfollowUser_Success(t *testing.T) {
	t.Log("US0007: Unfollow - sucesso ao deixar de seguir")
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)

	service := NewFollowService(followRepoMock, userRepoMock)

	followerID := uint(1)
	followedID := uint(2)

	followRepoMock.On("IsFollowing", followerID, followedID).Return(true)
	followRepoMock.On("Delete", followerID, followedID).Return(nil)

	err := service.UnfollowUser(followerID, followedID)

	assert.NoError(t, err)
	followRepoMock.AssertExpectations(t)
}

func TestGetFollowersCount_Success(t *testing.T) {
	t.Log("US0002: Contagem de seguidores do vendedor")
	followRepoMock := new(mocks.FollowRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)

	service := NewFollowService(followRepoMock, userRepoMock)

	userID := uint(1)
	user := &domain.User{
		UserID:   userID,
		UserName: "maria",
		UserType: domain.UserTypeSeller,
	}

	userRepoMock.On("FindByID", userID).Return(user, nil)
	followRepoMock.On("CountFollowers", userID).Return(int64(2))

	result, err := service.GetFollowersCount(userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result.UserID)
	assert.Equal(t, "maria", result.UserName)
	assert.Equal(t, int64(2), result.FollowersCount)
}

func TestGetFollowedList_OrderNameAsc(t *testing.T) {
	t.Log("US0004/US0008: Lista de vendedores seguidos (nome asc)")
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

	followed := []*domain.User{
		{UserID: 2, UserName: "ana", UserType: domain.UserTypeSeller},
		{UserID: 3, UserName: "carlos", UserType: domain.UserTypeSeller},
	}

	userRepoMock.On("FindByID", userID).Return(user, nil)
	followRepoMock.On("GetFollowed", userID, order).Return(followed, nil)

	result, err := service.GetFollowedList(userID, order)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.Followed))
	assert.Equal(t, "ana", result.Followed[0].UserName)
	assert.Equal(t, "carlos", result.Followed[1].UserName)

	userRepoMock.AssertExpectations(t)
	followRepoMock.AssertExpectations(t)
}

func TestGetFollowedList_OrderNameDesc(t *testing.T) {
	t.Log("US0004/US0008: Lista de vendedores seguidos (nome desc)")
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

	followed := []*domain.User{
		{UserID: 3, UserName: "carlos", UserType: domain.UserTypeSeller},
		{UserID: 2, UserName: "ana", UserType: domain.UserTypeSeller},
	}

	userRepoMock.On("FindByID", userID).Return(user, nil)
	followRepoMock.On("GetFollowed", userID, order).Return(followed, nil)

	result, err := service.GetFollowedList(userID, order)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.Followed))
	assert.Equal(t, "carlos", result.Followed[0].UserName)
	assert.Equal(t, "ana", result.Followed[1].UserName)

	userRepoMock.AssertExpectations(t)
	followRepoMock.AssertExpectations(t)
}
