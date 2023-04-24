package main

import "fmt"

func add(num1, num2 int) int {
	return num1 + num2
}

func subtraction(x, y int) int {
	return x - y

}

func main() {
	fmt.Println(add(23, 90))
	fmt.Println(subtraction(20, 13))
}
