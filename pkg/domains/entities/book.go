package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)



type Book struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    ISBN        string `json:"isbn"`
    PublishYear int    `json:"publishYear"`
    Genre       string `json:"genre"`
}






type BookReply struct {
	Data       *Book
	Collection []Book
	Stream      <-chan Book
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
