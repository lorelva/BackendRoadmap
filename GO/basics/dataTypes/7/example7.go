package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 8, 6
	var convert float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(convert)
	fmt.Println(x, y, convert)
}
