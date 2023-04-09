package main

import "fmt"

// it makes a double return
func swapFunc(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swapFunc("Valle", "Lorena")
	fmt.Println(a, b)
}
