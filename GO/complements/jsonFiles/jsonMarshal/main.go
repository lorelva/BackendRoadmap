package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// struct to json
func main() {

	//with tags,  but in salary will not be exported
	type employee1 struct {
		Name   string `json:"name,omitempty"`
		Age    int    `json:"age,omitempty"`
		salary int    `json:"salary,omitempty"`
	}

	//not using tags
	type employee2 struct {
		Name   string
		Age    int
		salary int
	}

	//Create values and initializes values of the struct for employee 1
	e1 := employee1{
		Name:   "Lorena",
		Age:    21,
		salary: 1000,
	}

	//Make the conversion
	j, err := json.Marshal(e1)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	fmt.Printf("employee1 JSON: %s\n", string(j))

	e2 := employee2{
		Name:   "Lyah",
		Age:    21,
		salary: 2000,
	}
	j, err = json.Marshal(e2)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	fmt.Printf("\nemployee2 JSON: %s\n", string(j))

}
