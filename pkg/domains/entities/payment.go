package entities

import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Payment struct {
	commons.FoundationEntity
	ID           uint       `json:"id"`
	OrderID       uint     `json:":order_id"`
	Amount       float64    `json:"amount"`
	Method       string     `json:"method"`
	Status       string      `json:"status"`
	Creditcard   Creditcard   `json:"creditcard"`
	
}

type PaymentReply struct {
	
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















