package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "48"
	num2 := "52"
	var x1, x2 int
	x1, _ = strconv.Atoi(num1)
	x2, _ = strconv.Atoi(num2)
	sum := x1 + x2
	fmt.Println("The sum of the two numbers are:", sum)
}
