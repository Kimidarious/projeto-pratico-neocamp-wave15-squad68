package domain

import (
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	FollowID uint         `gorm:"primaryKey;column:follow_id" json:"follow_id"`
	FollowerID uint       `gorm:"not null;column:follower_id;index" json:"follower_id"`
	FollowedID uint       `gorm:"not null;column:followed_id;index" json:"followed_id"`
	FollowDate time.Time  `gorm:"autoCreatedTime;column:follow_date" json:"follow_date"`

	Follower *User        `gorm:"foreignKey:FollowerID;references:UserID" json:"follower,omitempty"`
	Followed *User        `gorm:"foreignKey:FollowedID;references:UserID" json:"followed,omitempty"`
}

func (Follow) TableName() string {
	return "follows"
}

func (f *Follow) BeforeCreate(tx *gorm.DB) error {
	if f.FollowerID == f.FollowedID {
		return gorm.ErrInvalidData
	}
	return nil
}