package main

import "fmt"

/*
	Given an integer n, for every integer i <= n, the task is to print “FizzBuzz”
	if i is divisible by 3 and 5, “Fizz” if i is divisible by 3, “Buzz” if i is divisible
	by 5 or i (as a string) if none of the conditions are true.

*/

func main() {
	num := 30
	fizzBuzz(num)

}

func fizzBuzz(num int) {

	for i := 1; i < num; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("FizzBuzz")

		case i%3 == 0:
			fmt.Print("Fizz")

		case i%5 == 0:
			fmt.Print("Buzz")

		default:
			fmt.Print(i)

		}

		if i != num {
			fmt.Print(", ")
		}
	}
	fmt.Println()

}
