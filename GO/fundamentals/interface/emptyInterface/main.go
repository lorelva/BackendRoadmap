package main

import "fmt"

func main() {
	var x interface{}
	x = 5
	fmt.Printf("x=%v of type :%T\n", x, x)
	x = "Hello"
	fmt.Printf("x=%v of type :%T", x, x)

}
