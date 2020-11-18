package models

type Tokens struct {
	TokenId     int    `form:"token_id" json:"token_id"`
	Token       string `form:"token" json:"token" validate:"required"`
	Access      string `form:"access" json:"access" validate:"required"`
	RefreshType string `form:"refresh_type" json:"refresh_type" validate:"required"`
	CustomerId  string `form:"customer_id" json:"customer_id" validate:"required"`
}

