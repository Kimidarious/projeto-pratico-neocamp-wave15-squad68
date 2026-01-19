package request

type ProductRequest struct {
	ProductName string `json:"product_name" binding:"required,max=40"`
	Type        string `json:"type" binding:"required,max=15"`
	Brand       string `json:"brand,omitempty" binding:"max=50"`
	Color       string `json:"color,omitempty" binding:"max=20"`
	Notes       string `json:"notes,omitempty"`
}

type CreatePostRequest struct {
	UserID   uint           `json:"user_id" binding:"required,gt=0"`
	Date     string         `json:"date" binding:"required"`
	Product  ProductRequest `json:"product" binding:"required"`
	Category int            `json:"category" binding:"required,gt=0"`
	Price    float64        `json:"price" binding:"required,gt=0"`
}

type CreatePromoPostRequest struct {
	UserID   uint           `json:"user_id" binding:"required,gt=0"`
	Date     string         `json:"date" binding:"required"`
	Product  ProductRequest `json:"product" binding:"required"`
	Category int            `json:"category" binding:"required,gt=0"`
	Price    float64        `json:"price" binding:"required,gt=0"`
	HasPromo bool           `json:"has_promo"`
	Discount float64        `json:"discount" binding:"gte=0,lte=100"`
}