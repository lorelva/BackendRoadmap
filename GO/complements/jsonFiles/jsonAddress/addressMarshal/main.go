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
	address := Address{
		Street:  "123 Main St",
		City:    "New York",
		ZipCode: "10001",
	}

	jsonData, err := json.Marshal(address)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
