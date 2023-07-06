package entities



import(
	"github.com/msegeya56/ecommerce.go.module/pkg/commons"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/utils"
)



type Category struct {
	commons.FoundationEntity
	ID          int
	Name        string
	Description string
	Parent      *Category
	Subcategories []Category
	Products    []Product
}



type Checkout struct {
	commons.FoundationEntity
	ID         int
	Customer   Customer
	Products   []Product
	Total      float64
	Discount   float64
	PromoCode  string
    Completed  bool
}

type Customer struct {
	commons.FoundationEntity
	ID               int
	Username         string
	Email            string
	Password         string
	FullName         string
	Phone            string
	Address          string
	Orders           []Order
	Wishlist         []Product
	Reviews          []Review
	CreditCard       CreditCard
	ShippingAddress  string
	BillingAddress   string
	ProfilePicture   string
	AccountBalance   float64
	LastLogin        time.Time
	RegistrationDate time.Time
}





type Deposit struct {
	commosn.FoundationEntity
	ID        int
	Customer  Customer
	Amount    float64
	
}



type Invoice struct {
	commons.FoundationEntity
	ID         int
	Order      Order
	Amount     float64
	IssuedDate time.Time
	DueDate    time.Time
	Status     string
	Paid       bool
	Payment    Payment
	InvoiceURL string
}

ype Order struct {
	commons.FoundationEntity
	ID           int
	Customer     Customer
	Products     []Product
	TotalPrice   float64
	ShippingCost float64
	Discount     float64
	Status       string
	Payment      Payment
	ShippingInfo ShippingInfo
	Invoice      Invoice
	
}






type Payment struct {
	commons.FoundationEntity
	ID         int
	Order      entities.Order
	Amount     float64
	Method     string
	Status     string
	CardHolder string
	CardNumber string
	ExpiryDate string
	CVV        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}




type Product struct {
	commons.FoundationEntity
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
	Category    Category
	Tags        []string
	Reviews     []Review
	Ratings     []Rating
	Images      []string
}



type Receipt struct {
	common.FoundationEntity
	ID         int
	Invoice    Invoice
	Amount     float64
	Payment    entities.Payment
	
}






func (c *Customer) ToJSON() string {
	return utils.ToJSON(c)
}

func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}



func (r *Receipt) ToJSON() string {
	return utils.ToJSON(r)
}

func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}



func (ch**Checkout) ToJSON() string {
	return utils.ToJSON(ch)
}

func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}


func (p *Payment) ToJSON() string {
	return utils.ToJSON(p)
}

func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}


func (o *Order) ToJSON() string {
	return utils.ToJSON(o)
}



func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}


func (c *Customer) ToJSON() string {
	return utils.ToJSON(c)
}

func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}

func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}


func (p *Payment) ToJSON() string {
	return utils.ToJSON(p)
}

func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}


func (o *Order) ToJSON() string {
	return utils.ToJSON(o)
}



func FromJSON(jsonStr string, entity interface{}) error {
	return utils.FromJSON(jsonStr, entity)
}




func ( cr*CureditLimit) ToJSON() string {
	return utils.ToJSON(cr)
}

// func FromJSON(jsonStr string, entity interface{}) error {
// 	return utils.FromJSON(jsonStr, entity)
// }




func (i *Invoice) ToJSON() string {
	return utils.ToJSON(i)
}





func (d *Deposit) ToJSON() string {
	return utils.ToJSON(d)
}

func (ca *Category) ToJSON() string {
	return utils.ToJSON(ca)
}

























