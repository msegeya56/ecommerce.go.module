package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"


	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Creditcard struct {
    commons.FoundationEntity
    ID             uint   `gorm:"column:id;primaryKey;autoIncrement"`
    CardNumber     string `gorm:"column:card_number;type:varchar(25)"`
    CardholderName string `gorm:"column:card_holder_name;type:varchar(255)"`
    ExpirationDate string `gorm:"column:expiration_date;type:varchar(255)"`
    CVV            string `gorm:"column:cvv;type:varchar(10)"`
}






type CreditcardReply struct {
	Data       *Creditcard
	Collection []Creditcard
	streams    <-chan Creditcard
	Error      error
}


func (c *Creditcard) ToJson() string {
	jsonBytes, _ := json.Marshal(c)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (c *Creditcard) FromJson(data string) *Creditcard {
	err := json.Unmarshal([]byte(data), c)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return c
}

func (c *Creditcard) FromIOReadCloser(r io.ReadCloser) (*Creditcard, error) {

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

func (c *Creditcard) FromRequestBody(r *http.Request) (*Creditcard, error) {

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

func (c *Creditcard) FromResponseBody(r *http.Response) (*Creditcard, error) {

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