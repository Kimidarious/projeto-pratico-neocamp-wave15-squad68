package repository

import (
	"strings"

	"gorm.io/gorm"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type followRepositoryImpl struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepositoryImpl{db: db}
}

func (r *followRepositoryImpl) Create(follow *domain.Follow) error {
	return r.db.Create(follow).Error
}

func (r *followRepositoryImpl) Delete(followerID, followedID uint) error {
	return r.db.Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		Delete(&domain.Follow{}).Error
}

func (r *followRepositoryImpl) IsFollowing(followerID, followedID uint) bool {
	var count int64
	r.db.Model(&domain.Follow{}).
		Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		Count(&count)
	return count > 0
}

func (r *followRepositoryImpl) CountFollowers(userID uint) int64 {
	var count int64
	r.db.Model(&domain.Follow{}).
		Where("followed_id = ?", userID).
		Count(&count)
	return count
}

func (r *followRepositoryImpl) GetFollowers(userID uint, order string) ([]*domain.User, error) {
	var users []*domain.User

	query := r.db.Table("users").
		Joins("INNER JOIN follows ON users.user_id = follows.follower_id").
		Where("follows.followed_id = ?", userID)

	
	switch strings.ToLower(order) {
	case "name_asc":
		query = query.Order("users.user_name ASC")
	case "name_desc":
		query = query.Order("users.user_name DESC")
	default:
		query = query.Order("users.user_name ASC") // padrão
	}

	err := query.Find(&users).Error
	return users, err
}

func (r *followRepositoryImpl) GetFollowed(userID uint, order string) ([]*domain.User, error) {
	var users []*domain.User

	query := r.db.Table("users").
		Joins("INNER JOIN follows ON users.user_id = follows.followed_id").
		Where("follows.follower_id = ?", userID)

	
	switch strings.ToLower(order) {
	case "name_asc":
		query = query.Order("users.user_name ASC")
	case "name_desc":
		query = query.Order("users.user_name DESC")
	default:
		query = query.Order("users.user_name ASC") // padrão
	}

	err := query.Find(&users).Error
	return users, err
}