package request

type CreatePostRequest struct {
	UserID   uint           `json:"user_id" binding:"required,gt=0"`
	Date     string         `json:"date" binding:"required"`
	Product  ProductRequest `json:"product" binding:"required"`
	Category int           `json:"category" binding:"required,gt=0"`
	Price    float64        `json:"price" binding:"required,gt=0"`
}