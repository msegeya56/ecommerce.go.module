package entities

import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)

type Rating struct {
	commons.FoundationEntity
	ID          uint    `json:"id"`
    CustomerID uint      `json:"customer_id"`
    ProductID  uint      `json:"product_id"`
    OrderID    uint      `json:"order_id"`
    Rating     int       `json:"rating"`
    Comment    string      `json:"comment"`
}



type RatingReply struct {
	
	Data        *Rating
	collection  []Rating
	Stream      <-chan Rating
	Error       error
	ErrorStream <-chan error
}

func (ra *Rating) ToJson() string {
	jsonBytes, _ := json.Marshal(ra)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (ra *Rating) FromJson(data string) *Rating {
	err := json.Unmarshal([]byte(data), ra)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return ra
}

func (ra*Rating) FromIOReadCloser(r io.ReadCloser) (*Rating, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(ra)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return  ra,nil
}

func (ra*Rating) FromRequestBody(r *http.Request) (*Rating, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	// defer r.Body.close()

	for {
		err := decoder.Decode(ra)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ra, nil
}

func (ra *Rating) FromResponseBody(r *http.Response) (*Rating, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(ra)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ra, nil
}
