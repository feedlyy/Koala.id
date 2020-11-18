package models

type Orders struct {
	OrderId 			string `form:"order_id" json:"order_id"`
	CustomerId     	 	string `form:"customer_id" json:"customer_id" validate:"required"`
	OrderNumber     	string `form:"order_number" json:"order_number" validate:"required"`
	PaymentMethodId     string `form:"payment_method_id" json:"payment_method_id" validate:"required"`
}
