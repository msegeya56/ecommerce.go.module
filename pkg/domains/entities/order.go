package entities

import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
	"gorm.io/gorm"
)




type Order struct {
	gorm.Model
	// ID           uint         `gorm:"column:id;type:varchar;size:255"`
	Customer     Customer     `gorm:"column:customer;type:varchar;size:255"` 
	Products     []Product    `gorm:"column:products;type:varchar;size:255"`
	TotalPrice   float64        `gorm:"column:total;type:float64:val1;size:255"`
    Discount     float64         `gorm:"column:id;type:val1;size:255"`
	Status       string         `gorm:"column:status;type:string;size:255"`
	Payment      []Payment       `gorm:"column:payment;type:varchar;size:"`
	InvoiceID      uint          `gorm:"column:id;type:varchar;size:255"`
	
}

type OrderReply struct {
	commons.Foundation
	Data        *Order
	collection  []Order
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















