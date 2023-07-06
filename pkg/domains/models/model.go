package models



import()



type Category struct {
	common.FoundationEntity
	ID          int
	Name        string
	Description string
	Parent      *Category
	Subcategories []Category
	Products    []Product
}



type Checkout struct {
	common.FoundationEntity
	ID         int
	Customer   Customer
	Products   []Product
	Total      float64
	Discount   float64
	PromoCode  string
    Completed  bool
}

type Customer struct {
	common.FoundationEntity
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

func (c *Customer) toJSON() string {
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error marshaling customer to JSON:", err)
		return ""
	}
	return string(data)
}

func fromJSON(jsonStr string) (*Customer, error) {
	var c Customer
	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}




type Deposit struct {
	common.FoundationEntity
	ID        int
	Customer  Customer
	Amount    float64
	
}



type Invoice struct {
	common.FoundationEntity
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

func (o *Order) toJSON() string {
	data, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Error marshaling order to JSON:", err)
		return ""
	}
	return string(data)
}

func fromJSON(jsonStr string) (*Order, error) {
	var o Order
	err := json.Unmarshal([]byte(jsonStr), &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}






type Payment struct {
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

func (p *Payment) toJSON() string {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling payment to JSON:", err)
		return ""
	}
	return string(data)
}

func fromJSON(jsonStr string) (*Payment, error) {
	var p Payment
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}


type Product struct {
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

func (p *Product) toJSON() string {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling product to JSON:", err)
		return ""
	}
	return string(data)
}

func fromJSON(jsonStr string) (*Product, error) {
	var p Product
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}


type Receipt struct {
	common.FoundationEntity
	ID         int
	Invoice    Invoice
	Amount     float64
	Payment    entities.Payment
	
}

func (p *Receipt) toJSON() string {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling Receipt to JSON:", err)
		return ""
	}
	return string(data)
}

func fromJSON(jsonStr string) (*Receipt, error) {
	var p Receipt
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}














