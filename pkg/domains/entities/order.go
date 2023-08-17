package entities

import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"


	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Order struct {
	commons.FoundationEntity
	ID           uint         `json:"id"`
	CustomerID  uint         `json:":Customer_id"`
	Customer     Customer     `json:"customer"` 
	Products     []Product    `json:"products"`
	TotalPrice   float64        `json:"total_price"`
    Discount     float64         `json:"discount"`
	Status       string         `json:"status"`
	Payment      []Payment       `json:"payment"`
	InvoiceID      uint          `json:"invoice_id"`
	
}

type OrderReply struct {

	Data        *Order
	Collection  []Order
	Stream      <-chan Order
	Error       error
	ErrorStream <-chan error
}





func (o*Order) ToJson() string {
	jsonBytes, _ := json.Marshal(o)
	x:= fmt.Sprintf("%v", string(jsonBytes))

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














