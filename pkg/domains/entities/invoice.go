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
    ID           uint        `json:"id"`
    CustomerID   uint        `json:":customer_id"`
    Order        Order       `json:"order"`
    Amount       float64     `json:"amount"`
    IssuedDate   time.Time   `json:"issuedDate"`
    DueDate      time.Time   `json:"dueDate"`
    Status       string      `json:"status"`
    Paid         bool        `json:"paid"`
    Payment      Payment     `json:"paymentd"`
    InvoiceURL   string      `json:"invoiceURL"`


}






type INvoicetReply struct {
	
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


