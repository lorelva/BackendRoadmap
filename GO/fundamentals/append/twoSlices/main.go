package main

import "fmt"

func main() {

	//create two slices a&b
	a := []int{1, 4, 7}
	b := []int{5, 9}

	//Append one slice to another
	//NOTE: IMPORTANT TO ADD DOTS AT LAST
	//without (...)DOTS, code would attempt to append the slice as a whole, which is invalid.
	a = append(a, b...)

	//print the complement of the slice a and b
	//so... [1, 4, 7, 5,9]
	fmt.Println(a)

	/*viceversa
	b = append(b, a...)
	//so... [5,9, 1, 4, 7]
	fmt.Println(b)
	*/

}
