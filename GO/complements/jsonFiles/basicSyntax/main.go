package main

import (
	"encoding/json"
	"fmt"
)

// 1.- Create a JSON file (these case elements.json)

func main() {
	// 2.- Define data structure--> STRUCT in GO
	type Elements struct {
		//Declare each value with a capital letter, this for separating  json tags from struct
		Copper   int
		Krypton  int
		Platinum int
		Radon    int
	}

	// convert json data/string to go object/struct
	jsonStr := `{"copper": 29,"krypton": 36,"platinum": 78,"radon": 86}`
	fmt.Println(jsonStr)
	//go object type
	var toStruct Elements

	// call json library to unmarshall json data
	err := json.Unmarshal([]byte(jsonStr), &toStruct)
	//verify if the error is not null
	if err != nil {
		fmt.Println("Theres an error")

	}
	fmt.Println(toStruct)

}
