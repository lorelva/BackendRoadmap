package main

import "fmt"

func main() {

	//Create an array with the number 0 to 10
	//first way
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)

	//second way
	var numbers2 [10]int
	numbers2[0] = 1
	numbers2[1] = 2
	numbers2[2] = 3
	numbers2[3] = 4
	numbers2[4] = 5
	numbers2[5] = 6
	numbers2[6] = 7
	numbers2[7] = 8
	numbers2[8] = 9
	numbers2[9] = 10

	fmt.Println(numbers2)

	//just for fun
	for i := 1; i <= 10; i++ {
		fmt.Println(i)

	}

}
