package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)

type Category struct {
	commons.FoundationEntity
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	ParentID      uint       `json:"parent_id" gorm:"column:parent_id"`
	Parent        *Category  `json:"parent" gorm:"foreignKey:ParentID"`
	Subcategories []Category `json:"subcategories" gorm:"foreignKey:ParentID"`
	Products      []Product  `json:"products" gorm:"column:products;type:varchar;size:255"`
}

type CategoryReply struct {
	Data       *Category
	Collection []Category
	Stream     <-chan Category
	Error      error
}

func (ca *Category) ToJson() string {
	jsonBytes, _ := json.Marshal(ca)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (ca *Category) FromJson(data string) *Category {
	err := json.Unmarshal([]byte(data), ca)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return ca
}

func (ca *Category) FromIOReadCloser(r io.ReadCloser) (*Category, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(ca)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ca, nil
}

func (ca *Category) FromRequestBody(r *http.Request) (*Category, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(ca)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ca, nil
}

func (ca *Category) FromResponseBody(r *http.Response) (*Category, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(ca)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ca, nil
}
