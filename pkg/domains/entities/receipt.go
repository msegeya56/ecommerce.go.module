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
	commons.FoundationEntity
	CustomerID  uint        `gorm:"column:customer_id;type:varint;size:255"`
    OrderID       uint    `gorm:"column:order_id;type;varint:size255"`       // ID of the associated order
    Items         []ReceiptItem  `gorm:"column:items;type:varchar;size:255"`                                       
    Subtotal      float64         `gorm:"column:subtotal;type:float64;size:255"`
    Tax           float64         `gorm:"column:tax;type:float64;size:255"`
    Discount      float64         `gorm:"column:discount;type:float64;size:255"`
    TotalAmount   float64         `gorm:"column:total_amount;type:float64;size"`
    PaymentMethod string           `gorm:"column:payment_method;type:string;size:255"`
    TransactionDate time.Time       `gorm:"column:transactionDate;type:time.Time;size:255"`
}

type ReceiptItem struct {
    ProductName  string   `gorm:"column:product_name;type:string;size:255"`
    Quantity     int       `grorm:"column:quantity;type:255"`
     Subtotal     float64     `gorm:"column:subtotal;type:float64;size:255"`
}








type ReceiptReply struct {
	Data       *Receipt
	Collection []*Receipt
	streams    <-chan Receipt
	Error      error
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