package main

import "fmt"

func normalFunc(msg string) {
	fmt.Println(msg)
}

// with three arguments
func TripleArgument(x, y int, word string) {
	fmt.Println(x, y, word)
}

func main() {
	//or you can print sice here or call the name of the function and with the corresponding type of data
	normalFunc("Just Beginning... \n Yes I do\n ")
	TripleArgument(54, 67, "\nReject yourself")
}
