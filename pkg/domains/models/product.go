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




type Product struct {
	commons.Foundation
	ID          uint     `gorm:"column:id;type:varchar;size:255"`
	Name        string   `gorm:"column:name;type:varchar;size:255"`
	Description string   `gorm:"column:description;type:varchar;size:255"`
	Price       float64  `gorm:"column:price;type:varchar;size:255"`
	Stock       uint     `gorm:"column:stock;type:varchar;size:255"`
	Category    Category `gorm:"column:category;type:varchar;size:255"`
	Tags        []string `gorm:"column:tags;type:varchar;size:255"`
	Reviews     []Review `gorm:"column:reviews;type:varchar;size:255"`
	Ratings     []Rating `gorm:"column:ratings;type:varchar;size:255"`
	Images      []string `gorm:"column:images;type:varchar;size:255"`
}

type ProductReply struct {
	commons.Foundation
	Data        *entities.Product
	collection  []entities.Product
	Stream      <-chan entities.Product
	Error       error
	ErrorStream <-chan error
}




func (p*Product) ToJson() string {
	jsonBytes, _ := json.Marshal(p)
	x:= fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (p *Product) FromJson(data string) *Product {
	err := json.Unmarshal([]byte(data), p)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return p
}

func (p *Product) FromIOReadCloser(r io.ReadCloser) (*Product, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(p)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return p, nil
}

func (p *Product) FromRequestBody(r *http.Request) (*Product, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	// defer r.Body.close()

	for {
		err := decoder.Decode(p)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return p, nil
}

func (p *Product) FromResponseBody(r *http.Response) (*Product, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()


	for {
		err := decoder.Decode(p)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return p, nil
}















