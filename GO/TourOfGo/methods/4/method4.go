package main

import "fmt"

type Pentagon int

func (p Pentagon) Perimeter() int {
	return int(p) * 5
}

type Apotema int

func (a Apotema) Area() int {
	//return int(a) * Perimeter
}

func main() {
	p := Pentagon(7)
	fmt.Println(p.Perimeter())

	a := Apotema(4)
	fmt.Println(a.Area())
}
