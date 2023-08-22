package main

import "fmt"

// Slices are like references to arrays
func main() {
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	//separates the array like in two parts?
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	//porque las intercambia?
	b[0] = "HOLA"
	fmt.Println(a, b)
	fmt.Println(names)
}
