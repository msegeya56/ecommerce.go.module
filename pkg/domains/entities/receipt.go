package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)



type Receipt struct {
	commons.Foundation
	ID             uint           `json:"id"`
	CustomerID     uint           `json:"customer_id"`
	Customer       Customer       `gorm:"foreignKey:CustomerID" json:"customer"`
	OrderID        uint           `json:"order_id"`
	Order          Order          `gorm:"foreignKey:OrderID" json:"order"`
	Items          []ReceiptItem  `json:"items"`
	Subtotal       float64        `json:"subtotal"`
	Tax            float64        `json:"tax"`
	Discount       float64        `json:"discount"`
	TotalAmount    float64        `json:"total_amount"`
	PaymentMethod  string         `json:"payment_method"`
	TransactionDate time.Time      `json:"transactionDate"`
}






type ReceiptItem struct {
    ProductName  string   `json:"product_name"`
    Quantity     int       `json:"quantity"`
     Subtotal     float64     `json:"subtotal5"`
}








type ReceiptReply struct {
	Data       *Receipt
	Collection []Receipt
	Stream    <-chan Receipt
	Error      error
	ErrorStream <-chan error
}


func (re*Receipt) ToJson() string {
	jsonBytes, _ := json.Marshal(re)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func(re*Receipt) FromJson(data string) *Receipt {
	err := json.Unmarshal([]byte(data), re)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return re
}

func (re*Receipt) FromIOReadCloser(r io.ReadCloser) (*Receipt, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(re)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return re, nil
}

func (re*Receipt) FromRequestBody(r *http.Request) (*Receipt, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(re)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return re, nil
}

func (d*Receipt) FromResponseBody(r *http.Response) (*Receipt, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(d)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return d, nil
}