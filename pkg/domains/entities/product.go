package entities

import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)

type Product struct {
	commons.FoundationEntity
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       uint     `json:"stock"`
	Category    Category `json:"category"`
	Tags        []string `json:"tags"`
	Reviews     []Review `gorm:"foreignKey:ProductID" json:"reviews"`
	Ratings     []Rating `json:"ratings"`
	Images      []string `json:"images"`
}


type ProductReply struct {
	Data        *Product
	Collection  []Product
	Stream      <-chan Product
	Error       error
	ErrorStream <-chan error
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