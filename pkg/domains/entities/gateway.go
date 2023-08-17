package entities

import "github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"



type  Gateway struct {
	commons.FoundationEntity
	ID          uint    `json:"id"`
    Name        string `json:"name"`
    Description string  `json:"description"`
    APIKey      string   `json:"api_keye"`

	Product      ProductService
	Order        OrderService
	Gateway     GatewayService
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


type CustomerService struct {
	Customer []Customer
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







type CategoryService struct{
	Categories []Category
}


type DepositService struct{
	Deposits []Deposit
}


type GatewayService    struct{
	Gateways  []Gateway
}


type InvoiceService struct{
	Invoices  []Invoice
}












type CreditlimitService struct{
	Creditlimits  []Creditlimit
}





type GatewayReply struct {
	Data       *Gateway
	Collection []Gateway
	Streams    <-chan Gateway
	Error      error
}