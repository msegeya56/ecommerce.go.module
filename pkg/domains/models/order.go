package models

import (
	"encoding/json"
	"io"
	"time"

	"errors"
	"fmt"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)


type Order struct {
    commons.Foundation
    CustomerID   uint      `gorm:"column:customer_id" json:"customer_id"`
    Items        []OrderItem `gorm:"-" json:"items"`
    TotalAmount  float64   `gorm:"column:total_amount" json:"total_amount"`
    OrderDate    time.Time `gorm:"column:order_date" json:"order_date"`
    // Add other fields as needed
}

type OrderItem struct {
    ID         uint    `gorm:"primaryKey;column:id" json:"id"`
    OrderID    uint    `gorm:"column:order_id" json:"order_id"`
    ProductID  uint    `gorm:"column:product_id" json:"product_id"`
    Quantity   int     `gorm:"column:quantity" json:"quantity"`
    Price      float64 `gorm:"column:price" json:"price"`
    Product    Product `gorm:"foreignKey:ProductID" json:"product"`
}



type OrderReply struct {
	commons.Foundation
	Data        *entities.Order
	collection  []*entities.Order
	Stream      <-chan entities.Order
	Error       error
	ErrorStream <-chan error
}




func (o*Order) ToJson() string {
	jsonBytes, _ := json.Marshal(o)
	x:= fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (o *Order) FromJson(data string) *Order {
	err := json.Unmarshal([]byte(data), o)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return o
}

func (o *Order) FromIOReadCloser(r io.ReadCloser) (*Order, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(o)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return o, nil
}

func (o *Order) FromRequestBody(r *http.Request) (*Order, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	// defer r.Body.close()

	for {
		err := decoder.Decode(o)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return o, nil
}

func (o *Order) FromResponseBody(r *http.Response) (*Order, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()


	for {
		err := decoder.Decode(o)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return o, nil
}















