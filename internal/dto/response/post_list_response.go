package response

type ProductDTO struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Type        string  `json:"type"`
	Brand       string  `json:"brand,omitempty"`
	Color       string  `json:"color,omitempty"`
	Notes       string  `json:"notes,omitempty"`
}

type PostDTO struct {
	PostID   uint       `json:"post_id"`
	UserID   uint       `json:"user_id"`
	Date     string     `json:"date"`
	Product  ProductDTO `json:"product"`
	Category int        `json:"category"`
	Price    float64    `json:"price"`
	HasPromo bool       `json:"has_promo,omitempty"`
	Discount float64    `json:"discount,omitempty"`
}

type PostListResponse struct {
	UserID uint      `json:"user_id"`
	Posts  []PostDTO `json:"posts"`
}