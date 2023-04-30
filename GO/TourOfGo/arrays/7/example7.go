package main

import "fmt"

func main() {
	//Iterating over an array

	numbers := [8]int{34, 67, 89, 12, 10, 8, 678}
	var sum int = 0
	//or sum := int(0)

	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}

	//the numbers of the array
	fmt.Println("Numbers", numbers)
	//the sum of all the numbers
	fmt.Printf("Total  = %d\n", sum)
}
