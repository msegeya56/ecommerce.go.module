package utils

import (
	"encoding/json"
	"fmt"
)

func ToJSON(entity interface{}) string {
	data, err := json.Marshal(entity)
	if err != nil {
		fmt.Println("Error marshaling entity to JSON:", err)
		return ""
	}
	return string(data)
}




func ToJSON(model interface{}) string {
	data, err := json.Marshal(model)
	if err != nil {
		fmt.Println("Error marshaling model to JSON:", err)
		return ""
	}
	return string(data)
}







func FromJSON(jsonStr string, entity interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), entity)
	if err != nil {
		return err
	}
	return nil
}


unc FromJSON(jsonStr string, model interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), model)
	if err != nil {
		return err
	}
	return nil
}