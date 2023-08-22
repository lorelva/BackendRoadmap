package main

import "fmt"

//syntax of a function
/*  func nameFunction(parameters and data Type) (return values) {
	function body
}
*/

func add(x int, y int) int {
	//return directly the two parametrs x+y
	return x + y
}

// function continued
// Because it's the same data type it only requires declared it one time : x, y int
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(78, 10))
}
