package main

import "fmt"

// FOR CONTINUED--->The init and post statements are optional.
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//NOTE : WITHOUT SEMICOLONS IN THE FOR STATEMENT --> is like Go's "while",
	//because there is not while statement in GO
}
