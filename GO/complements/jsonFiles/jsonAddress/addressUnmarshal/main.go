package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	jsonStr := `{
        "street": "Av Emiliano Zapata",
        "city": "Cuernavaca",
        "zip_code": "62020"
    }`

	var address Address
	err := json.Unmarshal([]byte(jsonStr), &address)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	//access to struct information
	fmt.Printf("Street: %s\n City: %s\n   Zip Code: %s\n", address.Street, address.City, address.ZipCode)
}
