package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i < 5; i++ {
		numbers = append(numbers, i)
		fmt.Printf("Length: %d, Capacity: %d, Items: %v\n",
			len(numbers), cap(numbers), numbers)
	}

}
