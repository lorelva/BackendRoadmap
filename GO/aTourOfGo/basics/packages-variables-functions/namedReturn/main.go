package main

import "fmt"

// Naked return statements should be used only in short functions like this one
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//return x, y
	//it doesn't require it
	return
}

func main() {
	fmt.Println(split(17))
}
