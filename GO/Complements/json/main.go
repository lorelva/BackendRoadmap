package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Datos struct {
	FullName string `json:"full_name,omitempty"`
	Age      int    `json:"age,omitempty"`
}

func main() {
	content, err := os.ReadFile("./test.json")
	if err != nil {
		log.Fatal("hubo un error", err)
	}

	d := Datos{}
	m := make(map[string]interface{})

	json.Unmarshal(content, &m)

	err = json.Unmarshal(content, &d)
	if err != nil {
		log.Fatal("hubo un error", err)
	}

	// fmt.Println(m["full_name"])
	fmt.Println(d.FullName)
}
