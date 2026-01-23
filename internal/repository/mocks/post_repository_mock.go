package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type PostRepositoryMock struct {
	mock.Mock
}

func (m *PostRepositoryMock) Create(post *domain.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *PostRepositoryMock) FindByID(id uint) (*domain.Post, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Post), args.Error(1)
}

func (m *PostRepositoryMock) GetPostsByUser(userID uint) ([]*domain.Post, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Post), args.Error(1)
}

func (m *PostRepositoryMock) GetPostsByUsersInDateRange(userIDs []uint, startDate, endDate time.Time, order string) ([]*domain.Post, error) {
	args := m.Called(userIDs, startDate, endDate, order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Post), args.Error(1)
}

func (m *PostRepositoryMock) CountPromoPostsByUser(userID uint) int64 {
	args := m.Called(userID)
	return args.Get(0).(int64)
}

func (m *PostRepositoryMock) GetPromoPostsByUser(userID uint) ([]*domain.Post, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Post), args.Error(1)
}