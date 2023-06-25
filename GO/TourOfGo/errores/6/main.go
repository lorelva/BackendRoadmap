package main

import "fmt"

// create a function to divide two numbers
func division(num1, num2 int) error {

	// returns error, in this case, if num2 is 0 , cannot divde
	if num2 == 0 {
		return fmt.Errorf("%d / %d\n Cannot Divide Number by zero", num1, num2)
	}

	// returns the result
	return nil
}

func main() {

	err := division(10, 3)

	// error
	if err != nil {
		fmt.Printf("error: %s", err)

		// no error
	} else {
		fmt.Println("Valid Division")
	}
}
