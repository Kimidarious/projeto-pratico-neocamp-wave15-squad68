package repository

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type postRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepositoryImpl{db: db}
}

func (r *postRepositoryImpl) Create(post *domain.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepositoryImpl) FindByID(id uint) (*domain.Post, error) {
	var post domain.Post
	result := r.db.Preload("Product").Preload("User").First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, result.Error
	}
	return &post, nil
}

func (r *postRepositoryImpl) GetPostsByUser(userID uint) ([]*domain.Post, error) {
	var posts []*domain.Post
	err := r.db.Preload("Product").
		Where("user_id = ?", userID).
		Order("date DESC").
		Find(&posts).Error
	return posts, err
}

func (r *postRepositoryImpl) GetPostsByUsersInDateRange(userIDs []uint, startDate, endDate time.Time, order string) ([]*domain.Post, error) {
	var posts []*domain.Post

	query := r.db.Preload("Product").Preload("User").
		Where("user_id IN ?", userIDs).
		Where("date BETWEEN ? AND ?", startDate, endDate)

	
	switch strings.ToLower(order) {
	case "date_asc":
		query = query.Order("date ASC")
	case "date_desc":
		query = query.Order("date DESC")
	default:
		query = query.Order("date DESC")
	}

	err := query.Find(&posts).Error
	return posts, err
}

func (r *postRepositoryImpl) CountPromoPostsByUser(userID uint) int64 {
	var count int64
	r.db.Model(&domain.Post{}).
		Where("user_id = ? AND has_promo = ?", userID, true).
		Count(&count)
	return count
}

func (r *postRepositoryImpl) GetPromoPostsByUser(userID uint) ([]*domain.Post, error) {
	var posts []*domain.Post
	err := r.db.Preload("Product").
		Where("user_id = ? AND has_promo = ?", userID, true).
		Order("date DESC").
		Find(&posts).Error
	return posts, err
}
