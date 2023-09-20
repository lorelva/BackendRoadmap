package main

import "fmt"

func main() {

	num := 23

	if num >= 10 {
		fmt.Println("Num is more than 10")
		//nested if means an if inside the first if statement
		if num > 20 {
			fmt.Println("Num is also more than 20")
		}
	} else {
		fmt.Println("Num is less than 10")
	}

}
