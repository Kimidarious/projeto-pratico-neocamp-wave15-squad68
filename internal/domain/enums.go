package domain


type UserType string

const (
	UserTypeBuyer  UserType = "BUYER"
	UserTypeSeller UserType = "SELLER"
	UserTypeBoth   UserType = "BOTH"
)