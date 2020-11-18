package models

type OrderDetails struct {
	OrderDetailId string `form:"order_detail_id" json:"order_detail_id"`
	OrderId       string `form:"order_id" json:"order_id" validate:"required"`
	ProductId     string `form:"product_id" json:"product_id" validate:"required"`
	Qty           int    `form:"qty" json:"qty" validate:"required"`
	CreatedDate   string `form:"created_date" json:"created_date" validate:"required"`
}
