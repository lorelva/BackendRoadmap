package main

import "fmt"

func main() {

	//declare the slice as usual
	a := []int{10, 7}

	//Included two additional numbers so....  OUTPUT--> [10, 7, 3, 5]
	//SYNTAX -->slice = append(slice, element_1, element_2…, element_N)
	a = append(a, 3, 5)
	fmt.Println(a)

}
