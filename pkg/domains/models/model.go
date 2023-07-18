package models

import (
	"time"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
	"github.com/vmihailenco/msgpack/v5"
)

type JSONSerializableu interface {
	ToMsgpack() ([]byte, error)
	FromMsgpack(data []byte) error
}

type Category struct {
	commons.Foundation
	ID            uint
	Name          string
	Description   string
	Parent        *Category
	Subcategories []Category
	Products      []Product
}

type CategoryReply struct {
	commons.Foundation
	Data        *entities.Category
	Collection  []entities.Category
	Stream      <-chan entities.Category
	Error       error
	ErrorStream <-chan error
}

func (c *CategoryReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *CategoryReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}

type Checkout struct {
	commons.Foundation
	ID        uint
	Customer  Customer
	Products  []Product
	Total     float64
	Discount  float64
	PromoCode string
	Completed bool
}

type CheckoutReply struct {
	commons.Foundation
	Data        *entities.Checkout
	Collection  []entities.Checkout
	Stream      <-chan entities.Checkout
	Error       error
	ErrorStream <-chan error
}

func (ch *CheckoutReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(ch)
}

func (ch *CheckoutReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &ch)
}

type Creditlimit struct {
	commons.Foundation
	ID        uint
	Customer  Customer
	Products  []Product
	Total     float64
	Discount  float64
	PromoCode string
	Completed bool
}

type CreditlimitReply struct {
	commons.Foundation
	Data        *entities.Creditlimit
	Collection  []entities.Creditlimit
	Stream      <-chan entities.Creditlimit
	Error       error
	ErrorStream <-chan error
}

func (cl *CreditlimitReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(cl)
}

func (cl *CreditlimitReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &cl)
}

type Customer struct {
	commons.Foundation
	ID               uint        `json:"id"`
	Username         string      `json:"username"`
	Email            string      `json:"email"`
	Password         string      `json:"password"`
	FullName         string      `json:"fullName"`
	Phone            string      `json:"phone"`
	Address          string      `json:"address"`
	Orders           []Order     `json:"orders"`
	Wishlist         []Product   `json:"wishlist"`
	Reviews          []Review    `json:"reviews"`
	Creditcardd     Creditcard  `json:"creditCard"`
	ShippingAddress  string      `json:"shippingAddress"`
	BillingAddress   string      `json:"billingAddress"`
	ProfilePicture   string      `json:"profilePicture"`
	AccountBalance   float64     `json:"accountBalance"`
	LastLogin        time.Time   `json:"lastLogin"`
	
}

type CustomerReply struct {
	commons.Foundation
	Data        *entities.Customer
	Collection  []entities.Customer
	Stream      <-chan entities.Customer
	Error       error
	ErrorStream <-chan error
}

func (c *CustomerReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *CustomerReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}





type Deposit struct {
	commons.Foundation
	ID       uint          `json:"id"`
	Customer Customer     `json:"customer"`
	Amount   float64      `json:"a,oumt"`
}

type DepositReply struct {
	commons.Foundation
	Data        *entities.Deposit
	Collection  []entities.Deposit
	Stream      <-chan entities.Deposit
	Error       error
	ErrorStream <-chan error
}

func (d *DepositReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(d)
}

func (d *Deposit) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &d)
}

type Invoice struct {
               commons.Foundation
	ID           uint      `json:"id"`
	Order        Order     `json:"order"`
	Amount       float64   `json:"amount"`
	IssuedDate   time.Time `json:"issuedDate"`
	DueDate      time.Time `json:"dueDate"`
	Status       string    `json:"status"`
	Paid         bool      `json:"paid"`
	Payment      Payment   `json:"payment"`
	InvoiceURL   string    `json:"invoiceURL"`

}






type INvoicetReply struct {
	commons.Foundation
	Data        *entities.Invoice
	Collection  []entities.Invoice
	Stream      <-chan entities.Invoice
	Error       error
	ErrorStream <-chan error
}

func (i *INvoicetReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(i)
}

func (i *INvoicetReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &i)
}

type Order struct {
	commons.Foundation
	ID           uint      `json:"id"`
	Customer     Customer  `json:"customer"`
	Products     []Product `json:"products"`
	TotalPrice   float64   `json:"total_price"`
	ShippingCost float64   `json:"shipping_cost"`
	Discount     float64   `json:"discount"`
	Status       string    `json:"status"`
	Payment      []Payment `json:"payment"`
	Invoices     []Invoice `json:"invoices"`
}

type OrdertReply struct {
	commons.Foundation
	Data        *entities.Order
	Collection  []entities.Order
	Stream      <-chan entities.Order
	Error       error
	ErrorStream <-chan error
}

func (o *OrdertReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(o)
}

func (o *OrdertReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &o)
}

type Payment struct {
	commons.Foundation
	ID           uint       `json:"id"`
	Order        Order      `json:"order"`
	Amount       float64    `json:"amount"`
	Method       string     `json:"method"`
	Status       string     `json:"status"`
	Creditcard   Creditcard `json:"credit_card"`
}

type PaymentReply struct {
	commons.Foundation
	Data        *entities.Payment
	Collection  []entities.Payment
	Stream      <-chan entities.Payment
	Error       error
	ErrorStream <-chan error
}

func (pa *PaymentReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(pa)
}

func (pa PaymentReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &pa)
}

type Product struct {
	commons.Foundation
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       uint     `json:"stock"`
	Category    Category `json:"category"`
	Tags        []string `json:"tags"`
	Reviews     []Review `json:"reviews"`
	Ratings     []Rating `json:"ratings"`
	Images      []string `json:"images"`
}

type ProductReply struct {
	commons.Foundation
	Data        *entities.Product
	Collection  []entities.Product
	Stream      <-chan entities.Product
	Error       error
	ErrorStream <-chan error
}

func (pr *ProductReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(pr)
}

func (pr *ProductReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &pr)
}

type Receipt struct {
	commons.Foundation
	ID       uint     `json:"id"`
	Invoice  Invoice  `json:"invoice"`
	Amount   float64  `json:"amount"`
	Payments []Payment `json:"payments"`


}

type ReceiptReply struct {
	commons.Foundation
	Data        *entities.Receipt
	Collection  []entities.Receipt
	Stream      <-chan entities.Receipt
	Error       error
	ErrorStream <-chan error
}

func (r *Receipt) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(r)
}

func (r *Receipt) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &r)
}

func (rr *ReceiptReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(rr)
}

func (rr *ReceiptReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &rr)
}

type Review struct {
	commons.Foundation
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	ProductID uint      `json:"productId"`
	Rating    float64   `json:"rating"`
	Comment   string    `json:"comment"`
	Date      time.Time `json:"date"`


}

type ReviewReply struct {
	commons.Foundation
	Data        *entities.Review
	Collection  []entities.Review
	Stream      <-chan entities.Review
	Error       error
	ErrorStream <-chan error
}

func (re *ReviewReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(re)
}

func (re *ReviewReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &re)
}

type Creditcard struct {
	commons.Foundation
	CardNumber     string `json:"card_number"`
	CardholderName string `json:"cardholder_name"`
	ExpirationDate string `json:"expiration_date"`
	CVV            string `json:"cvv"`
}

type CreditcardReply struct {
	commons.Foundation
	Data        *entities.Creditcard
	Collection  []entities.Creditcard
	Stream      <-chan entities.Creditcard
	Error       error
	ErrorStream <-chan error
}

func (cc *CreditcardReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(cc)
}

func (cc *CreditcardReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &cc)
}

type Rating struct {
	commons.Foundation
	ProductID uint    `json:"productId"`
	Average   float64 `json:"average"`
	Count     uint    `json:"count"`
}


type RatingReply struct {
	commons.Foundation
	Data        *entities.Rating
	Collection  []entities.Rating
	Stream      <-chan entities.Rating
	Error       error
	ErrorStream <-chan error
}

func (ra *RatingReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(ra)
}

func (ra *RatingReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &ra)
}
