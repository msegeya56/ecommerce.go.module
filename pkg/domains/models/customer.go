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




type Customer struct {
	commons.Foundation
	Alias        string  `gorm:"column:alias" json:"alias,omitempty" gorm:"column:alias" dgraph:"customer.alias"`
 	ComplianceID string  `gorm:"column:compliance_id" json:"compliance_id,omitempty" gorm:"column:compliance_id" dgraph:"customer.complianceID"`
	Firstname    string  `gorm:"column:firstname" json:"firstname,omitempty" form:"first_name,omitempty" dgraph:"customer.firstname"`
    Middlename    string  `gorm:"column:middlename" json:"middlename,omitempty" form:"middlename,omitempty" dgraph:"customer.middlename"`
	Lastname     string  `gorm:"column:lastname" json:"lastname,omitempty" form:"lastname,omitempty" dgraph:"customer.lastname"`
	MobileNumber string  `gorm:"column:mobile_number" json:"mobile_number,omitempty" form:"mobile_number,omitempty" dgraph:"customer.mobilenumber"`
	Email        string  `gorm:"column:email" json:"email,omitempty" form:"email,omitempty" dgraph:"customer.email"`
	Address      string  `gorm:"column:address" json:"address,omitempty" form:"address,omitempty" dgraph:"customer.address"`
	Latitude     float64 `gorm:"column:latitude" json:"latitude,omitempty" form:"latitude,omitempty" dgraph:"customer.latitude"`
	Longitude    float64 `gorm:"column:longitude" json:"longitude,omitempty" form:"longitude,omitempty" dgraph:"customer.longitude"`
	Dob          string  `gorm:"column:dob" json:"dob,omitempty" form:"dob,omitempty" dgraph:"customer.dob"`
}


type CustomerReply struct {
	Data       *entities.Customer
	Stream     <-chan entities.Customer
	Collection []entities.Customer
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
			fmt.Println("Error decoding gorm:", err)
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
			fmt.Println("Error decoding gorm:", err)
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
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return c, nil
}