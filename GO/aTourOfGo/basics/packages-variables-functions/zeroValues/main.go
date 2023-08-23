package main

import "fmt"

// ZERO VALUES
// Variables that are declared without an explicit initial value will take for default : zero values.
func main() {
	//ZERO VALUES GIVEN FOR DEFAULT
	var i int // Numeric types --> 0
	var f float64
	var b bool   // Boolean types--> false
	var s string // String types--> " "

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
