package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Datos struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	d := Datos{
		Name: "Lorena",
		Age:  21,
	}

	content, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	jsonPrueba := ` { "name": "Lorena"  }  `
	content, _ = json.Marshal(jsonPrueba)
	fmt.Println(string(content))

	mapaPrueba := make(map[string]string)
	mapaPrueba["name"] = "Lorena"
	content, _ = json.Marshal(mapaPrueba)
	fmt.Println(string(content))

}
