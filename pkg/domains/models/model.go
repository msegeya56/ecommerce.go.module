package models

import (
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
	ProductID uint `json:"product_id,omitempty"`
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
	OrderID       uint `json:"order_id,omitempty"`
	DepositID     uint `json:"deposit_id,omitempty"`
	CreditlimitID uint `json:"credit_limit_id,omitempty"`
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
	DepositID uint `json:"deposit_id,omitempty"`
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
	CustomerID   uint `json:"customer_id,omitempty"`
	CreditCardID uint `json:"credit_card_id,omitempty"`
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
	Order Order
}

type INvoicetReply struct {
	commons.Foundation
	Data        *entities.Deposit
	Collection  []entities.Deposit
	Stream      <-chan entities.Deposit
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
	CustomerID uint `json:"customer_id,omitempty"`
	Products   []Product
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
	CustomerID   uint `json:"customer_id,omitempty"`
	OrderID      uint `json:"order_id,omitempty"`
	CreditCardID uint `json:"credit_card_id,omitempty"`
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
	CustomerID uint `json:"customer_id,omitempty"`
	CategoryID uint `json:"category_id,omitempty"`
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
	PaymentID uint `json:"payment_id,omitempty"`
	InvoiceID uint `json:"invoice_id,omitempty"`
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
	CustomerID uint `json:"customer_id,omitempty"`
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
	ProductID uint `json:"product_id,omitempty"`
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



