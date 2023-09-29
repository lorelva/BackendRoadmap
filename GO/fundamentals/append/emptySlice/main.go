package main

import "fmt"

func main() {
	//append to an empty slice:
	a := []int{}

	//Append returns the values included, because it was empty at the beginning
	a = append(a, 4, 6)
	fmt.Println(a)

}
