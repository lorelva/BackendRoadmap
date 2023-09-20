package main

import "fmt"

func main() {
	//Iterating over an array

	colors := [5]string{"Blue", "Gray", "Black", "Yellow", "Pink"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

}
