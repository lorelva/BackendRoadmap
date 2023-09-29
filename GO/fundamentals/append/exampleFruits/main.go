package main

import "fmt"

func main() {

	fruits := []string{"Watermelon, Apple"}
	moreFruits := []string{"Grapes, Orange"}

	fruits = append(fruits, moreFruits...)
	fmt.Println(fruits)

}
