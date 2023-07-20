package entities

import (
	"encoding/json"
	"io"
	"time"

	"errors"
	"fmt"
	"net/http"


	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)


type Invoice struct {
    commons.FoundationEntity
    ID           uint        `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
    CustomerID   uint        `gorm:"column:customer_id;type:varint;size:255"`
    Order        Order       `json:"order" gorm:"embedded"`
    Amount       float64     `json:"amount" gorm:"type:double;size:255"`
    IssuedDate   time.Time   `json:"issuedDate" gorm:"type:datetime"`
    DueDate      time.Time   `json:"dueDate" gorm:"type:datetime"`
    Status       string      `json:"status" gorm:"type:varchar(100)"`
    Paid         bool        `json:"paid"`
    Payment      Payment     `json:"payment" gorm:"embedded"`
    InvoiceURL   string      `json:"invoiceURL" gorm:"type:varchar(255)"`


}






type INvoicetReply struct {
	commons.Foundation
	Data        *Invoice
	Collection  []Invoice
	Stream      <-chan Invoice
	Error       error
	ErrorStream <-chan error
}



func (i*Invoice) ToJson() string {
	jsonBytes, _ := json.Marshal(i)
	x:= fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (i *Invoice) FromJson(data string) *Invoice {
	err := json.Unmarshal([]byte(data), i)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return i
}

func (i *Invoice) FromIOReadCloser(r io.ReadCloser) (*Invoice, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(i)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return i, nil
}

func (i*Invoice) FromRequestBody(r *http.Request) (*Invoice, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	// defer r.Body.close()

	for {
		err := decoder.Decode(i)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return i, nil
}

func (i *Invoice) FromResponseBody(r *http.Response) (*Invoice, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()


	for {
		err := decoder.Decode(i)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return i, nil
}


