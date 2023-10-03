package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"


	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)


type Customer struct {
    commons.FoundationEntity

    Firstname       string  `json:"firstname,omitempty" dgraph:"customer.firstname" xml:"firstname" gorm:"column:firstname" form:"firstname"`
    Lastname        string  `json:"lastname,omitempty" dgraph:"customer.lastname" xml:"lastname" gorm:"column:lastname" form:"lastname"`
    Middlename      string  `json:"middlename,omitempty" dgraph:"customer.middlename" xml:"middlename" gorm:"column:middlename" form:"middlename"`
    MobileNumber    string  `json:"mobile_number,omitempty" dgraph:"customer.mobileNumber" xml:"mobile_number" gorm:"column:mobile_number" form:"mobile_number"`
    Address         string  `json:"address,omitempty" dgraph:"customer.address" xml:"address" gorm:"column:address" form:"address"`
    Latitude        float64 `json:"latitude,omitempty" dgraph:"customer.latitude" xml:"latitude" gorm:"column:latitude" form:"latitude"`
    Longitude       float64 `json:"longitude,omitempty" dgraph:"customer.longitude" xml:"longitude" gorm:"column:longitude" form:"longitude"`
    Alias           string  `json:"alias,omitempty" dgraph:"customer.alias" xml:"alias" gorm:"column:alias" form:"alias"`
    Email           string  `json:"email,omitempty" dgraph:"customer.email" xml:"email" gorm:"column:email" form:"email"`
    Dob             string  `json:"dob,omitempty" dgraph:"customer.dob" xml:"dob" gorm:"column:dob" form:"dob"`
    Name            string  `json:"name,omitempty" dgraph:"customer.name" xml:"name" gorm:"column:name" form:"name"`
    ShippingAddress string  `json:"shippingAddress,omitempty" xml:"shippingAddress" form:"shippingAddress"`
    BillingAddress  string  `json:"billingAddress,omitempty" xml:"billingAddress" form:"billingAddress"`
    AccountBalance  float64 `json:"accountBalance,omitempty" xml:"accountBalance" form:"accountBalance"`
	Next            string  `json:"next,omitempty" dgraph:"customer.next" xml:"next" gorm:"column:next" form:"next"`
    Link            string  `json:"link,omitempty" dgraph:"customer.link" xml:"link" gorm:"column:link" form:"link"`
    Previous        string  `json:"previous,omitempty" dgraph:"customer.previous" xml:"previous" gorm:"column:previous" form:"previous"`
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