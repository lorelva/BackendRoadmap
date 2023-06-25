package main

import "fmt"

func main() {

	age := -14

	// create an error using Efforf()
	error := fmt.Errorf("%d is negative\n Theres no Age negative", age)

	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Println("Age: %d", age)
	}
}
