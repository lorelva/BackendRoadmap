package main

import (
	"fmt"
	"math"
)

// items of different type requires an explicit conversion
// This code can be the same as the code commit it:
func main() {

	var num1, num2 int = 5, 10
	var toFloat float64 = math.Sqrt(float64(num1 * num2))
	//var toFloat = math.Sqrt(float64(num1 * num2))
	var finalConversion uint = uint(toFloat)
	//var finalConversion  = uint(toFloat)
	fmt.Println(finalConversion)

	//call the function inference
	inference()

}

// TYPE INFERENCE:
func inference() {
	value := 46 // change me!
	//when you change the value in the printlf shows the real data type on what you assign
	fmt.Printf("Value is of type %T\n ", value)
}
