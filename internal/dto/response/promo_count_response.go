package response

type PromoCountResponse struct {
	UserID             uint   `json:"user_id"`
	UserName           string `json:"user_name"`
	PromoProductsCount int64  `json:"promo_products_count"`
}