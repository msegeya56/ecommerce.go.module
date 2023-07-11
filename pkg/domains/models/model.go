package models

import (
	

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
	"github.com/vmihailenco/msgpack/v5"
)

type JSONSerializableu interface{
	ToMsgpack() ([]byte, error)
	FromMsgpack(data []byte) error
}

type Category struct {
	commons.Foundation

} 



type  CategoryReply struct {
	commons.Foundation
	Data        *entities.Category
	Collection  []entities.Category
	Stream      <-chan entities.Category
	Error       error
	ErrorStream <-chan error
}




func (c *Category) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *Category) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}

func (c *CategoryReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *CategoryReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}





type Checkout struct {
	commons.Foundation

}

type  CheckoutReply struct {
	commons.Foundation
	Data        *entities.Checkout
	Collection  []entities.Checkout
	Stream      <-chan entities.Checkout
	Error       error
	ErrorStream <-chan error
}



func (ch *Checkout) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(ch)
}

func (ch *Checkout) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &ch)
}


func (ch *CheckoutReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(ch)
}

func (ch *CheckoutReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &ch)
}










type CreditLimit struct {
	commons.Foundation


}


type  CreditLimitReply struct {
	commons.Foundation
	Data        *entities.CreditLimit
	Collection  []entities.CreditLimit
	Stream      <-chan entities.CreditLimit
	Error       error
	ErrorStream <-chan error
}






func (cl *CreditLimit) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(cl)
}

func (cl *CreditLimit) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &cl)
}

func (cl *CreditLimitReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(cl)
}

func (cl *CreditLimitReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &cl)
}








type Customer struct {
	commons.Foundation
	
}


type  CustomerReply struct {
	commons.Foundation
	Data        *entities.Customer
	Collection  []entities.Customer
	Stream      <-chan entities.Customer
	Error       error
	ErrorStream <-chan error
}










func (c *Customer) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *Customer) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}



func (c *CustomerReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *CustomerReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}











type Deposit struct {
	commons.Foundation
}



type  DepositReply struct {
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
	Order      Order
	
}

type  INvoicetReply struct {
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
	ID          uint
	Customer     Customer
	Products     []Product
	
}


type  OrdertReply struct {
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
		Order      Order

}





type  PaymentReply struct {
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
	Category   Category
	
}



type  ProductReply struct {
	commons.Foundation
	Data        *entities.Product
	Collection  []entities.Product
	Stream      <-chan entities.Product
	Error       error
	ErrorStream <-chan error
}







func (p *Product) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(p)
}

func (p *Product) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &p)
}


func (pr *ProductReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(pr)
}

func (pr*ProductReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &pr)
}





type Receipt struct {
	commons.Foundation
	ID     uint
	Invoice Invoice
	
}

type  ReceiptReply struct {
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

}

type   ReviewReply struct {
	commons.Foundation
	Data        *entities.Review
	Collection  []entities.Review
	Stream      <-chan entities.Review
	Error       error
	ErrorStream <-chan error
}






func (r *Review) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(r)
}

func (r*Review) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &r)
}



func (re *ReviewReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(re)
}

func (re *ReviewReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &re)
}








type Rating struct {
	commons.Foundation
	ProductID uint
}

type   RatingReply struct {
	commons.Foundation
	Data        *entities.Rating
	Collection  []entities.Rating
	Stream      <-chan entities.Rating
	Error       error
	ErrorStream <-chan error
}






func (r *Rating) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(r)
}

func (r *Rating) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &r)
}



func ( ra *RatingReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(ra)
}

func ( ra *RatingReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &ra)
}








type CreditCard struct {
commons.Foundation

}


type   CreditCardReply struct {
	commons.Foundation
	Data        *entities.CreditCard
	Collection  []entities.CreditCard
	Stream      <-chan entities.CreditCard
	Error       error
	ErrorStream <-chan error
}



func (c *CreditCard) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *CreditCard) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}



func (ca *CreditCardReply) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(ca)
}

func (ca *CreditCardReply) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &ca)
}



