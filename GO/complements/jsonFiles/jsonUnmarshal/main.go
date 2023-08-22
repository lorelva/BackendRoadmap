package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee1 struct {
	Name   string `json:"n"`
	Age    int    `json:"a"`
	salary int    `json:"s"`
}

type employee2 struct {
	Name   string
	Age    int
	salary int
}

func main() {
	e1Json := `{"n":"Lyah","a":21}`

	var e1Converted employee1
	err := json.Unmarshal([]byte(e1Json), &e1Converted)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Printf("employee1 Struct: %#v\n", e1Converted)

	e2Json := `{"Name":"Lyah","Age":21}`
	var e2Converted employee2
	err = json.Unmarshal([]byte(e2Json), &e2Converted)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Printf("\nemployee2 Struct: %#v\n", e2Converted)
}
