package main

import "fmt"

// L -->LISKOV SUBSTITUTION PRINCIPLE (LSP)
type Animal struct {
	Name string
}

func (a Animal) MakeSound() {
	fmt.Println("Animal sound")
}
