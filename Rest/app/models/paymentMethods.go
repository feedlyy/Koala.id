package models

type PaymentMethods struct {
	PaymentMethodId string `form:"payment_method_id" json:"payment_method_id"`
	MethodName      string `form:"method_name" json:"method_name" validate:"required"`
	CreatedDate     string `form:"created_date" json:"created_date" validate:"required"`
}

