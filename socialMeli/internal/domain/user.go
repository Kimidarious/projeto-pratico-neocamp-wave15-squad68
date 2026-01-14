package domain

import (
	"time"
)

type UserType string

const (
	UserTypeBuyer  UserType = "BUYER"
	UserTypeSeller UserType = "SELLER"
	UserTypeBoth   UserType = "BOTH"
)

type User struct {
	UserID    uint      `gorm:"primaryKey;column:user_id" json:"user_id"`
	UserName  string    `gorm:"size:15;not null;unique;column:user_name" json:"user_name"`
	UserType  UserType  `gorm:"type:varchar(10);not null;column:user_type" json:"user_type"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
