package main

import "fmt"

// zero values
// Variables declared without an explicit initial value are given their zero value.
func main() {
	var i int
	var f float64
	var b bool
	var s string

	//separate ways
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", b)
	fmt.Printf("%q\n", s)

	//all in one line code
	fmt.Printf("%v %v %v %q\n\n", i, f, b, s)
}
