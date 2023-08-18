package models

import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)

type Order struct {
	commons.Foundation
	ID         uint      `gorm:"column:id;type:varchar;size:255"`
	Customer   Customer  `gorm:"foreignkey:CustomerID"`
	CustomerID uint      `gorm:"column:customer_id;type:varchar;size:255"`
	Products   []Product `gorm:"column:products;type:varchar;size:255"`
	TotalPrice float64   `gorm:"column:total_price;type:float64;size:255"`
	Discount   float64   `gorm:"column:discount;type:float64;size:255"`
	Status     string    `gorm:"column:status;type:string;size:255"`
	Payment    []Payment `gorm:"column:payment;type:varchar;size:255"`
	InvoiceID  uint      `gorm:"column:invoice_id;type:varchar;size:255"`
}

type OrderReply struct {
	commons.Foundation
	Data        *entities.Order
	Collection  []entities.Order
	Stream      <-chan entities.Order
	Error       error
	ErrorStream <-chan error
}

func (o *Order) ToJson() string {
	jsonBytes, _ := json.Marshal(o)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (o *Order) FromJson(data string) *Order {
	err := json.Unmarshal([]byte(data), o)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return o
}

func (o *Order) FromIOReadCloser(r io.ReadCloser) (*Order, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(o)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return o, nil
}

func (o *Order) FromRequestBody(r *http.Request) (*Order, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	// defer r.Body.close()

	for {
		err := decoder.Decode(o)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return o, nil
}

func (o *Order) FromResponseBody(r *http.Response) (*Order, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(o)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return o, nil
}
