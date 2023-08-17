package entities







import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)




type Review struct {
	commons.FoundationEntity
	CustomerID uint    `json:"customer_id"`
	ProductID  uint    `json:"product_id"`
	OrderID    uint    `json:"order_id"`
	Rating     float64 `json:"rating"`
	Comment    string  `json:"comment"`
}




type ReviewReply struct {
	Data        *Review
	Collection  []Review
	Stream      <-chan Review
	Error       error
	ErrorStream <-chan error
}




func (re*Review) ToJson() string {
	jsonBytes, _ := json.Marshal(re)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (re *Review) FromJson(data string) *Review {
	err := json.Unmarshal([]byte(data), re)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return re
}

func (re*Review) FromIOReadCloser(r io.ReadCloser) (*Review, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(re)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return  re,nil
}

func (re*Review) FromRequestBody(r *http.Request) (*Review, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	// defer r.Body.close()

	for {
		err := decoder.Decode(re)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return re, nil
}

func (re *Review) FromResponseBody(r *http.Response) (*Review, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(re)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return re, nil
}
