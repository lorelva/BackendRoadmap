package main

import "fmt"

func main() {

	numList := []int{2, 5, 6, 34, 90, 3}
	number := 3
	result := findNumber(numList, number)
	fmt.Println(result)

}

func findNumber(numbers []int, number int) bool {

	for i := 0; i < numbers[i]; i++ {
		if number == numbers[i] {
			return true
		}
	}
	return false
}
