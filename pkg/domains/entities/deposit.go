package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"gorm.io/gorm"
	// "github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Deposit struct {
	gorm.Model
	// ID            uint      `json:"id" gorm:"id"` 
	CustomerID  uint        `gorm:"column:customer_id;type:varint;size:255"`
	Amount   float64        `gorm:"column:amount;type:val1;size:255"`
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