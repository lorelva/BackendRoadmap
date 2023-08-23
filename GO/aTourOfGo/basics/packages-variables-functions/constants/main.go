package main

import "fmt"

// Constants are declared like variables, but with the const keyword.
// Cannot be declared using the := syntax.
const Pi = 3.14

func main() {
	const World = "Merhaba"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
