package models

import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)

type Rating struct {
	commons.FoundationEntity
	ID          uint    `gorm:"column:id;type:varchar;size:255"`
    CustomerID uint      `gorm:"column:customer_id;type:varchar;type:varchar;size:255"`
    ProductID  uint      `gorm:"column:product_id;type:varchar;size:255"`
    OrderID    uint      `gorm:"column:order_id;type:varchar;size:255"`
    Rating     int       `gorm:"column:rating;type:varchar;size:255"`
    Comment    string      `gorm:"column:comment;type:varchar;size:255"`
}



type RatingReply struct {
	commons.FoundationEntity
	Data        *entities.Rating
	collection  []entities.Rating
	Stream      <-chan entities.Rating
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
