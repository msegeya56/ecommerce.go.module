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
	Name        string   `gorm:"column:name;type:varchar(255)" json:"name"`
	Description string   `gorm:"column:description;type:varchar(255)" json:"description"`
	Price       float64  `gorm:"column:price" json:"price"`
	Stock       uint     `gorm:"column:stock" json:"stock"`
	CategoryID  uint     `gorm:"column:category_id" json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
}

type Tag struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	Name      string `gorm:"column:name;type:varchar(255)" json:"name"`
	ProductID uint   `gorm:"column:product_id" json:"product_id"`
}

type Review struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	Text      string `gorm:"column:text;type:text" json:"text"`
	ProductID uint   `gorm:"column:product_id" json:"product_id"`
}

type Rating struct {
	ID        uint    `gorm:"primaryKey;column:id" json:"id"`
	Score     float64 `gorm:"column:score" json:"score"`
	ProductID uint    `gorm:"column:product_id" json:"product_id"`
}

type Image struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	URL       string `gorm:"column:url;type:varchar(255)" json:"url"`
	ProductID uint   `gorm:"column:product_id" json:"product_id"`
}

func (p *Product) ToJson() string {
	jsonBytes, _ := json.Marshal(p)
	x := fmt.Sprintf("%v", string(jsonBytes))

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
