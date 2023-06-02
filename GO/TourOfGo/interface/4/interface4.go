package main

import (
	"fmt"
	"math"
)

// Una interfaz nos ayuda a agrupar funciones de metodos
type Figuras interface {
	Area() interface{}
	// Perimetro() float64
}

type Cuadrado int64

func (c *Cuadrado) Area() interface{} {
	return int64(*c) * int64(*c)
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t *Triangulo) Area() interface{} {
	area := t.Altura * t.Base
	return area
}

type Circulo float64

func (c *Circulo) Area() interface{} {
	area := math.Pi * math.Pow(float64(*c), 2)
	return area
}

type Rectangulo struct {
	Base   int64
	Altura int64
}

func (r *Rectangulo) Area() interface{} {
	area := r.Base * r.Altura
	return area
}

// la importancia de usar(guardar) interface sin saber el tipo de dato en especifico
func imprimir(valor interface{}) {
	fmt.Printf("Tipo de datos: %T el valor es: %v\n", valor, valor)
}

func main() {
	var (
		opc int = 0
		f   Figuras
	)

	c := Cuadrado(10)
	t := Triangulo{
		Base:   12.78,
		Altura: 18.89,
	}
	circ := Circulo(5.7)

	r := Rectangulo{
		Base:   12,
		Altura: 19,
	}

	for {

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

		resultado := f.Area()
		imprimir(resultado)
	}
}
