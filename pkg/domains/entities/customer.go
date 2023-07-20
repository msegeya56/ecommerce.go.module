package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Customer struct {
	commons.FoundationEntity
	ID               uint        `json:"id"`
	Username         string      `json:"username"`
	Email            string      `json:"email"`
	Password         string      `json:"password"`
	FullName         string      `json:"fullName"`
	Phone            string      `json:"phone"`
	Address          string      `json:"address"`
	Orders           []Order     `json:"orders"`
	Wishlist         []Product   `json:"wishlist"`
	Reviews          []Review    `json:"reviews"`
	Creditcardd     Creditcard  `json:"creditCard"`
	ShippingAddress  string      `json:"shippingAddress"`
	BillingAddress   string      `json:"billingAddress"`
	ProfilePicture   string      `json:"profilePicture"`
	AccountBalance   float64     `json:"accountBalance"`
	LastLogin        time.Time   `json:"lastLogin"`
}








type CustomerReply struct {
	Data       *Customer
	Collection []Customer
	streams    <-chan Customer
	Error      error
}


func (c *Customer) ToJson() string {
	jsonBytes, _ := json.Marshal(c)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (c *Customer) FromJson(data string) *Customer {
	err := json.Unmarshal([]byte(data), c)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return c
}

func (c *Customer) FromIOReadCloser(r io.ReadCloser) (*Customer, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(c)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return c, nil
}

func (c *Customer) FromRequestBody(r *http.Request) (*Customer, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(c)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return c, nil
}

func (c *Customer) FromResponseBody(r *http.Response) (*Customer, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(c)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return c, nil
}