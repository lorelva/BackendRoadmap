package main

import (
	"fmt"
	"math"
)

// If with a short statement
func pow(num1, num2, limit float64) float64 {
	// validate if the power of num1 and 2 is less than the limit
	//return the value if the condition was true
	if value := math.Pow(num1, num2); value < limit {
		return value
		//else : show why the validation before was above the limit
	} else {
		fmt.Printf("%g >= %g\n", value, limit)
	}
	return limit
}

func main() {
	fmt.Println(
		pow(5, 2, 10),
		pow(2, 8, 20),
	)
}
