package main

import "fmt"

func main() {
	//the number will be 4321
	//last digit plus one --> 1+1 = 2
	digits := []int{4, 3, 2, 1}
	//output --> [4 3 2 2]
	result := plusOne(digits)
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	/*to find the last digit in the array and then to plus one
	lastDigit := digits[len(digits)-1]
	lastDigit += 1
	fmt.Println(lastDigit)
	*/
	//rest digits in the array
	digits[len(digits)-1] = digits[len(digits)-1] + 1
	//allDigits := append([]int{digits}, lastDigit)
	return digits

}
