package response

type UserDTO struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}

type FollowersListResponse struct {
	UserID    uint      `json:"user_id"`
	UserName  string    `json:"user_name"`
	Followers []UserDTO `json:"followers"`
}

type FollowedListResponse struct {
	UserID   uint      `json:"user_id"`
	UserName string    `json:"user_name"`
	Followed []UserDTO `json:"followed"`
}