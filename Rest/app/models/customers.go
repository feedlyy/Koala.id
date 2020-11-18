package models

import "time"

var test = []string{"081285398772"}

type Customers struct {
	CustomerId  	string `form:"customer_id" json:"customer_id"`
	CustomerName  	string `form:"customer_name" json:"customer_name" validate:"required"`
	Email       	string `form:"email" json:"email" validate:"required,email"`
	PhoneNumber     string `form:"phone_number" json:"phone_number" validate:"required"`
	Dob         	time.Time `form:"dob" json:"dob" validate:"required"`
	Sex         	string `form:"sex" json:"sex" validate:"required"`
	Salt        	string `form:"salt" json:"salt"`
	Password    	string `form:"password" json:"password" validate:"required"`
	CreatedDate 	time.Time `form:"created_date" json:"created_date"`
}