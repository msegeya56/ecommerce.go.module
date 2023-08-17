package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Deposit struct {
	commons.FoundationEntity
	ID          uint        `json:"id"`   
	CustomerID  uint        `json:"customer_id"`
	Amount   float64        `json:"amount"`
}








type DepositReply struct {
	Data       *Deposit
	Collection []Deposit
	streams    <-chan Deposit
	Error      error
}


func (d*Deposit) ToJson() string {
	jsonBytes, _ := json.Marshal(d)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func(d*Deposit) FromJson(data string) *Deposit {
	err := json.Unmarshal([]byte(data), d)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return d
}

func (d*Deposit) FromIOReadCloser(r io.ReadCloser) (*Deposit, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

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

func (d*Deposit) FromRequestBody(r *http.Request) (*Deposit, error) {

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

func (d*Deposit) FromResponseBody(r *http.Response) (*Deposit, error) {

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