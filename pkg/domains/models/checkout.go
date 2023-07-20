package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"


	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Checkout struct {
	commons.Foundation
	ID          uint        `gorm:"column:id;type:varint;size:255"`   
	CustomerID  uint        `gorm:"column:customer_id;type:varint;size:255"`
     Products  []Product     `gorm:"column:products;type:varint;size:255"`
	Total     float64        `gorm:"column:total;type:float64:varint:size:255"`
	Discount  float64        `gorm:"column:discount;type:float64:varint:size :f255"`
	PromoCode string         `gorm:"column:promo_code;type:string;varchar:size;size:255"`
	Completed bool           `gorm:"column:completed;type:text;varchar:size;size:255"`
}








type CheckoutReply struct {
	Data       *entities.Checkout
	Collection []entities.Checkout
	streams    <-chan entities.Checkout
	Error      error
}


func (ch*Checkout) ToJson() string {
	jsonBytes, _ := json.Marshal(ch)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func(ch*Checkout) FromJson(data string) *Checkout {
	err := json.Unmarshal([]byte(data), ch)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return ch
}

func (ch*Checkout) FromIOReadCloser(r io.ReadCloser) (*Checkout, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(ch)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ch, nil
}

func (ch*Checkout) FromRequestBody(r *http.Request) (*Checkout, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(ch)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ch, nil
}

func (ch*Checkout) FromResponseBody(r *http.Response) (*Checkout, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(ch)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ch, nil
}