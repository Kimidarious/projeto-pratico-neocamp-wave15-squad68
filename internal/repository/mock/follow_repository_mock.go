package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type FollowRepositoryMock struct {
	mock.Mock
}

func (m *FollowRepositoryMock) Create(follow *domain.Follow) error {
	args := m.Called(follow)
	return args.Error(0)
}

func (m *FollowRepositoryMock) Delete(followerID, followedID uint) error {
	args := m.Called(followerID, followedID)
	return args.Error(0)
}

func (m *FollowRepositoryMock) IsFollowing(followerID, followedID uint) bool {
	args := m.Called(followerID, followedID)
	return args.Bool(0)
}

func (m *FollowRepositoryMock) CountFollowers(userID uint) int64 {
	args := m.Called(userID)
	return args.Get(0).(int64)
}

func (m *FollowRepositoryMock) GetFollowers(userID uint, order string) ([]*domain.User, error) {
	args := m.Called(userID, order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *FollowRepositoryMock) GetFollowed(userID uint, order string) ([]*domain.User, error) {
	args := m.Called(userID, order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}