package models

import (
	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	// "github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
	"gorm.io/gorm"
)




	type GatewayReply struct{
		Data     *entities.Gateway
		Collection []entities.Gateway
		Streams   <-chan entities.Gateway
		Error       error 
	}	





type  Gateway struct {
	gorm.Model
    Name        string `gorm:"column:name;type:varchar;size:255"`
    Description string  `gorm:"column:description;type:varchar;size:255"`
    APIKey      string   `gorm:"column:api_key;type:varchar;size:255"`
    Customer     CustomerService
	Product      ProductService
	Payment      PaymentService
   
    Receipt         ReceiptService


	// Add more microservices as needed
}


type CustomerService struct {
	Customer []Customer
}

type ProductService struct {
	Products []Product
}

type CheckoutService struct {
	CheckoutItems []Checkout
}

type PaymentService struct {
	Payments []Payment
}



type ReceiptService struct{
	Receipts []Receipt
}







type GatewayService    struct{
	Gateways  []Gateway
}












