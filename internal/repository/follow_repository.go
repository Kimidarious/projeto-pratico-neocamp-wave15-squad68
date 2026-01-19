package repository

import (
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type FollowRepository interface {
	Create(follow *domain.Follow) error
	Delete(followerID, followedID uint) error
	IsFollowing(followerID, followedID uint) bool
	CountFollowers(userID uint) int64
	GetFollowers(userID uint, order string) ([]*domain.User, error)
	GetFollowed(userID uint, order string) ([]*domain.User, error)
}