package main

import "fmt"

// TRAVERSING A STRING -----Iterates through the string with a loop.
// All characters in sequence.
func main() {
	str := "Lorena"
	for i, char := range str {
		fmt.Printf("%d, Character: %c\n", i, char)
	}
}
