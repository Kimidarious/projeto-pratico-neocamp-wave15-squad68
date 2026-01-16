package domain

import ("time")

type Post struct {
	PostID uint          `gorm:"primaryKey;column:post_id" json:"post_id"`
	UserID uint 		 `gorm:"not null;column:user_id;index" json:"user_id"`
	ProductID uint		 `gorm:"not null;column:product_id" json:"product_id"`
	Date time.Time		 `gorm:"type:date;not null;index" json:"date"`
	Category int		 `gorm:"not null" json:"category"`
	Price float64		 `gorm:"type:decimal(10,2);not null" json:"price"`
	HasPromo bool 		 `gorm:"default:false;column:has_promo;index" json:"has_promo"`
	Discount float64	 `gorm:"type:decimal(5,2);default:0" json:"discount,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`

	User *User           `gorm:"foreignKey:UserID;references:UserID" json:"user,omitempty"`
	Product *Product	 `gorm:"foreignKey:ProductID;references:ProductID" json:"product,omitempty"`
}

func (Post) TableName() string {
	return "posts"
}