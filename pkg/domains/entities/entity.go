package entities

import (
	"time"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
	"github.com/vmihailenco/msgpack/v5"
)

type JSONSerializableu interface {
	ToMsgpack() ([]byte, error)
	FromMsgpack(data []byte) error
}

type Category struct {
	commons.FoundationEntity
	ID            uint
	Name          string
	Description   string
	Parent        *Category
	Subcategories []Category
	Products      []Product
}

func (c *Category) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *Category) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}

type Checkout struct {
	commons.FoundationEntity
	ID        uint
	Customer  Customer
	Products  []Product
	Total     float64
	Discount  float64
	PromoCode string
	Completed bool
}

func (ch *Checkout) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(ch)
}

func (ch *Checkout) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &ch)
}

type Creditlimit struct {
	commons.FoundationEntity
	Customer Customer
	Limit    float64
	Used     float64
}

func (cl *Creditlimit) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(cl)
}

func (cl *Creditlimit) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &cl)
}

type Customer struct {
	commons.FoundationEntity
	ID               uint
	Username         string
	Email            string
	Password         string
	FullName         string
	Phone            string
	Address          string
	Orders           []Order
	Wishlist         []Product
	Reviews          []Review
	Creditcardd      Creditcard
	ShippingAddress  string
	BillingAddress   string
	ProfilePicture   string
	AccountBalance   float64
	LastLogin        time.Time
	RegistrationDate time.Time
}

func (c *Customer) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *Customer) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}

type Deposit struct {
	commons.FoundationEntity
	ID       uint
	Customer Customer
	Amount   float64
}

func (d *Deposit) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(d)
}

func (d *Deposit) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &d)
}

type Invoice struct {
	commons.FoundationEntity
	ID   uint    `json:"id"`
	OrderID     uint    `json:"order_id"`
	Amount      float64 `json:"amount"`
	IssuedDate  string  `json:"issued_date"`
	DueDate     string  `json:"due_date"`
	Status      string  `json:"status"`
	Paid        bool    `json:"paid"`
	PaymentID   uint    `json:"payment_id"`
	InvoiceURL  string  `json:"invoice_url"`
}



func (i *Invoice) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(i)
}

func (i *Invoice) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &i)
}

type Order struct {
	commons.FoundationEntity
	ID           uint
	Customer     Customer
	Products     []Product
	TotalPrice   float64
	ShippingCost float64
	Discount     float64
	Status       string
	Payment      []Payment
	Invoice      []Invoice
}

func (o *Order) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(o)
}

func (o *Order) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &o)
}

type Payment struct {
	commons.FoundationEntity
	ID          uint
	Order       Order
	Amount      float64
	Method      string
	Status      string
	Creditcardd Creditcard
}

func (p *Payment) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(p)
}

func (p *Payment) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &p)
}

type Product struct {
	commons.FoundationEntity
	ID          uint
	Name        string
	Description string
	Price       float64
	Stock       uint
	Category    Category
	Tags        []string
	Reviews     []Review
	Ratings     []Rating
	Images      []string
}

func (p *Product) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(p)
}

func (p *Product) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &p)
}

type Receipt struct {
	commons.FoundationEntity
	ID       uint
	Invoice  Invoice
	Amount   float64
	Payments []Payment
}

func (r *Receipt) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(r)
}

func (r *Receipt) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &r)
}

type Review struct {
	commons.FoundationEntity
	ID        uint
	UserID    uint
	ProductID uint
	Rating    float64
	Comment   string
	Date      time.Time
}

func (r *Review) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(r)
}

func (r *Review) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &r)
}

type Rating struct {
	ProductID uint
	Average   float64
	Count     uint
}

func (r *Rating) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(r)
}

func (r *Rating) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &r)
}

type Creditcard struct {
	commons.FoundationEntity
	CardNumber     string
	CardholderName string
	ExpirationDate string
	CVV            string
}

func (c *Creditcard) ToMsgpack() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *Creditcard) FromMsgpack(data []byte) error {
	return msgpack.Unmarshal(data, &c)
}
