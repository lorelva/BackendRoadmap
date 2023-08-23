package main

import (
	"fmt"
	"math/cmplx"
)

// Variable declarations may be "factored" into blocks into var statement
var (
	forString  string     = "Hello"
	forBool    bool       = false
	forFloat   float32    = 23.56
	forInt     uint64     = 1<<64 - 1
	forComplex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	// %T prints the types of the variables declared before
	fmt.Printf("Type: %T -->Value: %v\n", forString, forString)
	fmt.Printf("Type: %T -->Value: %v\n", forBool, forBool)
	fmt.Printf("Type: %T -->Value: %v\n", forFloat, forFloat)
	fmt.Printf("Type: %T -->Value: %v\n", forInt, forInt)
	fmt.Printf("Type: %T -->Value: %v\n", forComplex, forComplex)
}
