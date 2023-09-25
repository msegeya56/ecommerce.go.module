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
        FirstName    string `json:"firstname" gorm:"column:firstname" graphql:"firstname"`
        LastName     string `json:"lastname" gorm:"column:lastname" graphql:"lastname"`
        MiddleName   string `json:"middlename" gorm:"column:middlename" graphql:"middlename"`
        MobileNumber string `json:"mobile_number" gorm:"column:mobile_number" graphql:"mobile_number"`
        Latitude     float64 `json:"latitude" gorm:"column:latitude" graphql:"latitude"`
        Longitude    float64 `json:"longitude" gorm:"column:longitude" graphql:"longitude"`
        Alias        string `json:"alias" gorm:"column:alias" graphql:"alias"`
        Email        string `json:"email" gorm:"column:email" graphql:"email"`
        Dob          string `json:"dob" gorm:"column:dob" graphql:"dob"`
        Name         string `json:"name" gorm:"column:name" graphql:"name"`
        Next         string `json:"next" gorm:"column:next" graphql:"next"`
        Link         string `json:"link" gorm:"column:link" graphql:"link"`
        Previous     string `json:"previous" gorm:"column:previous" graphql:"previous"`
    
}


// type CustomerReply struct {
// 	Data       *entities.Customer
// 	Stream     <-chan entities.Customer
// 	Collection []entities.Customer
// 	Error      error
// }






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
