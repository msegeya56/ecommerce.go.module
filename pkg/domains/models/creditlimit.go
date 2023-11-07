package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"gorm.io/gorm"
)


type Creditlimit struct {
    gorm.Model 
    CustomerID     uint    `gorm:"column:customer_id;type:varint;size:255"`
    MaxLimit       float64 `gorm:"column:max_limit;type:double;size:255"`
    CurrentBalance float64 `gorm:"column:current_balance;type:double;size:255"`
}







type CreditlimittReply struct {
	Data       *entities.Creditlimit
	Collection []entities.Creditlimit
	streams    <-chan entities.Creditlimit
	Error      error
}


func (cl*Creditlimit) ToJson() string {
	jsonBytes, _ := json.Marshal(cl)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func(cl*Creditlimit) FromJson(data string) *Creditlimit {
	err := json.Unmarshal([]byte(data), cl)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return cl
}

func (cl*Creditlimit) FromIOReadCloser(r io.ReadCloser) (*Creditlimit, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(cl)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return cl, nil
}

func (ch*Creditlimit) FromRequestBody(r *http.Request) (*Creditlimit, error) {

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

func (cl*Creditlimit) FromResponseBody(r *http.Response) (*Creditlimit, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(cl)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return cl, nil
}
