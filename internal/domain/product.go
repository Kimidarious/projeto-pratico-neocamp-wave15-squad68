package domain

import (
	"time"
)

type Product struct {
	ProductID uint        `gorm:"primaryKey;column:product_id" json:"product_id"`
	ProductName string    `gorm:"size:40;not null;column:product_name" json:"product_name"`
	Type string			  `gorm:"size:15;not null" json:"type"`
	Brand string		  `gorm:"size:50" json:"brand,omitempty"`
	Color string		  `gorm:"size:20" json:"color,omitempty"`
	Notes string		  `gorm:"type:text" json:"notes,omitempty"`
	CreatedAt time.Time	  `gorm:"autoCreateTime" json:"created_at"`
}

func (Product) TableName() string {
	return "products"
}