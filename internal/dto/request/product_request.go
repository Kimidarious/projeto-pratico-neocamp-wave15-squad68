package request

type ProductRequest struct {
	ProductName string `json:"product_name" binding:"required,max=40"`
	Type        string `json:"type" binding:"required,max=15"`
	Brand       string `json:"brand,omitempty" binding:"max=50"`
	Color       string `json:"color,omitempty" binding:"max=20"`
	Notes       string `json:"notes,omitempty"`
}