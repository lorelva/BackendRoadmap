package main

import "fmt"

func main() {

	primes := [10]int{2, 3, 5, 7, 11, 13, 15, 17, 19, 21}
	var slice1 []int = primes[3:6]

	fmt.Println(slice1)
}
