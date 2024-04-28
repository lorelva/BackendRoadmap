package main

import (
	"fmt"

	"github.com/lorelva/BackendRoadmap/test/BasicStructure/message"
	"github.com/lorelva/BackendRoadmap/test/BasicStructure/str"
	"github.com/lorelva/BackendRoadmap/test/BasicStructure/suma"
)

func main() {
	result := suma.Add(1, 1)

	if result == 3 {
		fmt.Println("correcto")
	} else {
		fmt.Println("incorrecto")
	}

	result2, err := message.Message("Lorena")
	if result2 == "Lorena" && err != nil {
		fmt.Println("correcto")
	} else {
		fmt.Println("incorrecto")
	}

	result3, err := str.ConvertString("234")
	if result3 != 234 && err != nil {
		fmt.Println("correcto")
	} else {
		fmt.Println("incorrecto")
	}
}
