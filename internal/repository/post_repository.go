package repository

import (
	"time"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type PostRepository interface {
	Create(post *domain.Post) error
	FindByID(id uint) (*domain.Post, error)
	GetPostsByUser(userID uint) ([]*domain.Post, error)
	GetPostsByUsersInDateRange(userIDs []uint, startDate, endDate time.Time, order string) ([]*domain.Post, error)
	CountPromoPostsByUser(userID uint) int64
}