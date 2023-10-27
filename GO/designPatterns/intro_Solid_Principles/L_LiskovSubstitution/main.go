package main

import "fmt"

// L -->LISKOV SUBSTITUTION PRINCIPLE (LSP)

// STATEMENT: objects of a superclass should be replaceable with objects of a subclass ,
// without affecting the correctness of the program
// OBJECTIVE: demonstrates inheritance in Go

// EXAMPLE -->Animals and their sounds
type Animal struct {
	Name string
}

func (a Animal) MakeSound() {
	fmt.Println("Animal sound")
}

// Adding a first animal
type Bird struct {
	Animal
}

func (b Bird) MakeSound() {
	fmt.Println("Chirp chirp")
}

// adding a second animal
type Cow struct {
	Animal
}

func (c Cow) MakeSound() {
	fmt.Println("MUuu Muuuu")
}

type AnimalBehavior interface {
	MakeSound()
}

// Function to work with base class (Animal) or any subclass (Bird in this case)
func MakeSound(ab AnimalBehavior) {
	ab.MakeSound()
}

func main() {
	a := Animal{}
	b := Bird{}
	c := Cow{}
	MakeSound(a)
	MakeSound(b)
	MakeSound(c)
}
