package main

import "fmt"

// Nil slices : The zero value of a slice is nil.
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
