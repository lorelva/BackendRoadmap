package main

import "fmt"

func main() {
	result, err := Divide(3, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	//or just ignore the error
	/*
		result, _ := Divide(3, 0)
		fmt.Println(result)
	*/
}

func Divide(num1, num2 int) (int, error) {
	//if the divisor(num2) is zero, returns 0 and an error as result
	if num2 == 0 {
		return 0, fmt.Errorf("the divisor cannot be zero")
	}
	//if it's not zero the divisor, return the result of the division
	//and  nil error value
	return num1 / num2, nil

}
