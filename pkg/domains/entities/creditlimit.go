package entities



import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/msegeya56/ecommerce.go.module/pkg/tools/commons"
)


type Creditlimit struct {
    commons.Foundation
    ID             uint    `json:"id"`
    CustomerID     uint    `json:"customer_id"`
    MaxLimit       float64 `json:"max_limit"`
    CurrentBalance float64 `json:"current_balance"`
}







type CReditlimittReply struct {
	Data       *Creditlimit
	Collection []Creditlimit
	Stream    <-chan Creditlimit
	Error      error
	ErrorStream <-chan error
}


func (cl*Creditlimit) ToJson() string {
	jsonBytes, _ := json.Marshal(cl)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func(cl*Creditlimit) FromJson(data string) *Creditlimit {
	err := json.Unmarshal([]byte(data), cl)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return cl
}

func (cl*Creditlimit) FromIOReadCloser(r io.ReadCloser) (*Creditlimit, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(cl)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return cl, nil
}

func (ch*Creditlimit) FromRequestBody(r *http.Request) (*Creditlimit, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(ch)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return ch, nil
}

func (cl*Creditlimit) FromResponseBody(r *http.Response) (*Creditlimit, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(cl)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding JsON:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return cl, nil
}
