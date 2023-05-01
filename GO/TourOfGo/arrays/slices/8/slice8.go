package main

import "fmt"

func main() {
	slice1 := make([]int, 5)
	printSlice("a", slice1)

	slice2 := make([]int, 0, 5)
	printSlice("b", slice2)

	slice3 := slice2[:2]
	printSlice("c", slice3)

	slice4 := slice3[2:5]
	printSlice("d", slice4)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
