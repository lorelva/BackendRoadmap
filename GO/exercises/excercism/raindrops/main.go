package main

import "fmt"

/*
RULES FOR RAINDROPS
	IF the number:
	has 3 as a factor, add 'Pling' to the result.
	has 5 as a factor, add 'Plang' to the result.
	has 7 as a factor, add 'Plong' to the result.
	does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
*/

func main() {
	number := 800
	result := Convert(number)
	fmt.Println(result)

}

// validate the factors of the input number
func Convert(number int) string {

	var output string = ""
	if number%3 == 0 {
		output += "Ping"
	}
	if number%5 == 0 {
		output += "Plang"
	}
	if number%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output += fmt.Sprintf("%d", number)
	}

	return output
}
