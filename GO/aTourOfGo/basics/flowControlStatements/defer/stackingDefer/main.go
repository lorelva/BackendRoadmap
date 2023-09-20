package main

import "fmt"

func main() {

	//It will be executed first
	fmt.Println("\nCOUNTING....")

	//EXectued at third
	//deferred calls are executed in last-in-first-out order.
	//It means the reverse count 10...9..8..7...6........0
	for i := 0; i < 10; i++ {
		defer fmt.Println("Number:", i)

	}
	//it will be executed in second place
	fmt.Println("DONE")

}
