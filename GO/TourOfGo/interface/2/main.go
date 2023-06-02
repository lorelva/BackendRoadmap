package main

import (
	"fmt"
	"math"
)

type Figuras interface {
	Area()
	// Perimetro() float64
}

type Cuadrado float64

func (c *Cuadrado) Area() {

	area := float64(*c) * float64(*c)
	fmt.Println(area)
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t *Triangulo) Area() {
	area := t.Altura * t.Base
	fmt.Println(area)
}

type Circulo float64

func (c *Circulo) Area() {
	area := math.Pi * math.Pow(float64(*c), 2)
	fmt.Println(area)
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (r *Rectangulo) Area() {
	area := r.Base * r.Altura
	fmt.Println(area)
}

func main() {
	var (
		opc int = 0
		f   Figuras
	)

	c := Cuadrado(10.2)
	t := Triangulo{
		Base:   12.78,
		Altura: 18.89,
	}
	circ := Circulo(5.7)

	r := Rectangulo{
		Base:   12.67,
		Altura: 10.78,
	}

	fmt.Println("1.-Cuadrado")
	fmt.Println("2.-Triangulo")
	fmt.Println("3.-Circulo")
	fmt.Println("4.-Rectangulo")

	fmt.Scanf("%d", &opc)

	switch opc {
	case 1:
		f = &c

	case 2:
		f = &t

	case 3:
		f = &circ

	case 4:
		f = &r

	}
	f.Area()

}
