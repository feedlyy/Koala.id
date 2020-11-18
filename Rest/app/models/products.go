package models

type Products struct {
	ProductId   string `form:"product_id" json:"product_id"`
	ProductName string `form:"product_name" json:"product_name" validate:"required"`
	BasicPrice  int    `form:"basic_price" json:"basic_price" validate:"required"`
	CreatedDate string `form:"created_date" json:"created_date" validate:"required"`
}
