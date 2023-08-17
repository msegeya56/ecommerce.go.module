package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Customer struct {
	commons.Foundation
	ID               uint        `gorm:"column:id;type:varchar;size:255"`
	Username         string      `gorm:"column:username;type:varchar;size:255"`
	Email           string       `gorm:"column:email;type:varchar;size:255"`
	Password         string      `gorm:"column:password;type:varchar;size:255"`
	FullName         string      `gorm:"column:fullName;type:varchar;size:255"`
	Phone            string      `gorm:"column:phone;type:varchar;size:255"`
	Address          string      `gorm:"column:address;type:varchar;size:255"`
	Orders           []Order     `gorm:"column:orders;type:varchar;size:255"`
	Wishlist         []Product   `gorm:"column:wishlist;type:varchar;size:255"`
	Reviews          []Review    `gorm:"column:reviews;type:varchar;size:255"`
	Creditcardd     Creditcard  `gorm:"column:creditCard;type:varchar;size:255"`
	ProfilePicture   string      `gorm:"column:profilePicture;type:varchar;size:255"`
	AccountBalance   float64     `gorm:"column:accountBalance;type:varchar;size:255"`
	LastLogin        time.Time   `gorm:"column:lastLogin;type:varchar;size:255"`
}








type CustomerReply struct {
	Data       *entities.Customer
	Collection []entities.Customer
	streams    <-chan entities.Customer
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