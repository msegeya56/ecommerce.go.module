package models

import (
	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)


type Gateway struct {
	commons.Foundation
	ID           uint      `gorm:"column:id;type:varchar;size:255"`
	Name         string     `gorm:"column:name:varchar;type:varchar;size:255"`
	Description  string     `gorm:"column:description:varchar;type:varch:varchar;size:255"`
	APIKey       string      `gorm:"column:api_key:varchar;type:varchar;size:255"`

	Product      ProductService
	Order        OrderService
	Customer     CustomerService
	Category     CategoryService
	Checkout    CheckoutService
	Payment      PaymentService
    Rating        RatingService
	Review      RatingService
    Creditcard      CreditcardService
	Deposit         DepositService
	Creditlimit     CreditlimitService
	INvoice          InvoiceService
    Receipt         ReceiptService


	// Add more microservices as needed
}




	type GatewayReply struct{
		Data     *entities.Gateway
		Collection []entities.Gateway
		Streams   <-chan entities.Gateway
		error       error 
	}	





type CustomerService struct {
	Customers []Customer
}

type ProductService struct {
	Products []Product
}

type OrderService struct {
	Orders []Order
}

type CheckoutService struct {
	CheckoutItems []Checkout
}

type PaymentService struct {
	Payments []Payment
}

type CreditcardService struct {
	Creditcards []Creditcard
}

type RatingService struct {
	
	Ratings []Rating
}





type ReviewService struct{
	Reviews []Review
}

type ReceiptService struct{
	Receipts []Receipt
}

type InvoiceService struct{
	Invoice []Invoice

}




type CategoryService struct{
	Categories []Category
}

type DepositService struct{
	Deposits []Deposit
}


type  CreditlimitService struct{
	Creditlimits []Creditlimit
}