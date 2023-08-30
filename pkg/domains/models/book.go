package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
)





type Book struct {
	
    ID          uint   `gorm:"primaryKey"`
    Title       string `gorm:"column:title;not null"`
    Author      string `gorm:"column:author;not null"`
    ISBN        string `gorm:"column:isbn;not null"`
    PublishYear int    `gorm:"column:publish_year;not null"`
    Genre       string `gorm:"column:genre;not null"`
}


type BookReply struct {
	Data       *entities.Book
	Collection []entities.Book
	Stream      <-chan entities.Book
	Error      error
	ErrorStream <-chan error
}


func (b*Book) ToJson() string {
	jsonBytes, _ := json.Marshal(b)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func(b*Book) FromJson(data string) *Book {
	err := json.Unmarshal([]byte(data), b)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return b
}

func (b*Book) FromIOReadCloser(r io.ReadCloser) (*Book, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(b)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return b, nil
}

func (b*Book) FromRequestBody(r *http.Request) (*Book, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(b)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return b, nil
}

func (b*Book) FromResponseBody(r *http.Response) (*Book, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(b)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return b, nil
}


