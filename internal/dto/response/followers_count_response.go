package response

type FollowersCountResponse struct {
	UserID         uint   `json:"user_id"`
	UserName       string `json:"user_name"`
	FollowersCount int64  `json:"followers_count"`
}