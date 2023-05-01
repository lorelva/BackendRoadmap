package main

import "fmt"

// Slice defaults
// The default is zero for the low bound and the length of the slice for the high bound.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	//why the number to the position 2? instead of 1
	s = s[1:]
	fmt.Println(s)

}
