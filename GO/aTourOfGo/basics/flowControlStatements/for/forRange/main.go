package main

import "fmt"

// The use of range is to iterate over an array and print both the indexes and the values at each
func main() {

	//Declare an array with the elements: fruits
	fruits := []string{"Strawberry", "Grape", "Peach"}

	//i --> the position of each fruit
	//element --> store the value(friut in this case)
	//range friuts --> for iterate the array before
	for i, element := range fruits {
		fmt.Printf("%v\t%v\n", i, element)
	}
}
