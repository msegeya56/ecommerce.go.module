package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"

)




type Checkout struct {
	commons.FoundationEntity
	CustomerID  uint        `json:"customer_id"`
     Products  []Product     `json:"products"`
	Total     float64        `json:"total"`
	Discount  float64        `json:"discount"`
	PromoCode string         `json:"promo_code"`
	Completed bool           `json:"completed:"`
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