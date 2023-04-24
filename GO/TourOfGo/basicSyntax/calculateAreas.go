package main

import (
	"fmt"
	"math"
)

// calculate area of a circle
func circleArea(rad float64) float64 {
	return math.Pow(rad, 2) * math.Pi
}

// calculate area of a rectangle
func rectangleArea(base, height int) {
	rectArea := base * height
	fmt.Println("The rectangle area is:\n", rectArea)
}

func main() {
	//print the value
	fmt.Println(" The circle area is:\n", circleArea(6))
	//in this case only call the name of the function and the parameters
	rectangleArea(23, 6)
}
