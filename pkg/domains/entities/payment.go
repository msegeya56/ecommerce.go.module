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




type Payment struct {
	gorm.Model
	// ID           uint       `gorm:"column:id;type:varchar;size:255"`
	OrderID       uint     `gorm:"column:order_id;type:varchar;size:255"`
	Amount       float64    `gorm:"column:amount;type:varchar;size:255"`
	Method       string     `gorm:"column:method;type:varchar;size:255"`
	Status       string      `gorm:"column:status;type:varchar;size:255"`
	Creditcard   Creditcard   `gorm:"column:credit;type:varchar;size:255"`
	
}

type PaymentReply struct {
	commons.Foundation
	Data        *Payment
	collection  []Payment
	Stream      <-chan Payment
	Error       error
	ErrorStream <-chan error
}




func (pa*Payment) ToJson() string {
	jsonBytes, _ := json.Marshal(pa)
	x:= fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (pa *Payment) FromJson(data string) *Payment {
	err := json.Unmarshal([]byte(data), pa)



	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return pa
}

func (pa*Payment) FromIOReadCloser(r io.ReadCloser) (*Payment, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(pa)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return pa, nil
}

func (pa *Payment) FromRequestBody(r *http.Request) (*Payment, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	// defer r.Body.close()

	for {
		err := decoder.Decode(pa)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return pa, nil
}

func (pa *Payment) FromResponseBody(r *http.Response) (*Payment, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()


	for {
		err := decoder.Decode(pa)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return pa, nil
}















