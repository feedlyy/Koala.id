package models

type Tokens struct {
	TokenId     string `form:"token_id" json:"token_id"`
	Token       string `form:"token" json:"token"`
	RefreshType string `form:"refresh_type" json:"refresh_type"`
	CustomerId  string `form:"customer_id" json:"customer_id"`
}

