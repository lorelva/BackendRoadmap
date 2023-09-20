package main

import (
	"fmt"
	"math"
)

// create data custom
type Circle struct {
	radius float32
}

// create method for area
func (c Circle) Area() float32 {
	return (math.Pi) * (c.radius * c.radius)
}

// create method for perimeter
func (c Circle) Perimeter() float32 {
	return (2 * math.Pi) * (c.radius)
}

// create method to change the radius value
func (c *Circle) ChangeValue(r float32) {
	c.radius = r
}
func main() {
	c := Circle{
		radius: 6.7,
	}

	fmt.Println(c.radius, c.Area())
	fmt.Println(c.radius, c.Perimeter())

	//change the value and get new area and perimeter
	c.ChangeValue(10.8)
	fmt.Println(c.radius, c.Area(), c.Perimeter())
}
