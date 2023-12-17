package main

import "fmt"

func main() {
	numbers := []int{2, 5, 18, 4}
	sum := sumNumbers(numbers)
	fmt.Println(sum)
}

// optimized code
func sumNumbers(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + sumNumbers(numbers[1:])

}
