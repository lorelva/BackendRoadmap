package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// file .json to struct
// struct  that is recquired to calculate BMI: weight and height
type Person struct {
	FullName string  `json:"full_name,omitempty"`
	Age      int     `json:"age,omitempty"`
	Gender   string  `json:"gender,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Height   float64 `json:"height,omitempty"`
}

//Method CalculateBMI : body mass index (BMI) = IMC

func (p Person) Bmi() float32 {
	return (float32(p.Weight) / (float32(p.Height) * float32(p.Height)))
}

func main() {

	content, err := os.ReadFile("./person.json")
	if err != nil {
		log.Fatal("there was an error", err)
	}

	p := Person{}
	json.Unmarshal(content, &p)

	err = json.Unmarshal(content, &p)
	if err != nil {
		log.Fatal("there was an error", err)
	}

	fmt.Println("The IMC or BMI is", p.Bmi())

}
