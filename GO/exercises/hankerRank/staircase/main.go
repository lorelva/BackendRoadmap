package main

import (
	"fmt"
)

func main() {

	var n int32 = 6
	staircase(n)
	//fmt.Println(n)
}

func staircase(n int32) {
	for i := 1; i <= int(n); i++ {
		for j := 1; j <= int(n)-i; j++ {
			fmt.Print(" ")
		}
		for m := 1; m <= i; m++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}
