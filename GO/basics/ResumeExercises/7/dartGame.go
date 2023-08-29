package main

import (
	"math"
)

func calculatePoints(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)

	if distance > 10 {
		return 0
	} else if distance > 5 {
		return 1
	} else if distance > 1 {
		return 5
	} else {
		return 10
	}
}

func main() {
	// Example usage
	x := 5.0
	y := 8.0
	points := calculatePoints(x, y)
	println("Earned points:", points)
}
