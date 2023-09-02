package main

import "fmt"

func main() {

	numberLength := 30

	//fmt.Println(numbers, numberLength)
	showNumbers := (arrayNumbers(numberLength))

	fmt.Println(showNumbers)
}

func arrayNumbers(numberLength int) []int {

	arrNumbers := make([]int, numberLength)

	for i := 0; i < numberLength; i++ {

		arrNumbers[i] += i + 1

	}
	return arrNumbers

}
