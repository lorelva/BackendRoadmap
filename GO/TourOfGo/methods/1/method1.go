package main

import "fmt"

///EXTERNAL EXAMPLES

// 1.-Crear tipo de dato custom
// 2.-type {Nombre del tipo} {tipo de dato[map,int,int32, float64, func....]}
type Cuadrado int

// Crear metodo
// func ({nombre de variable} {Nombre del tipo}) nombreFuncion().......
func (c Cuadrado) Area() int {
	return int(c) * 4
}

type Triangulo float64

func (t Triangulo) Area() float64 {
	return float64(t) / 2
}

func main() {
	// Declaracion del tipo custom en una variable llamada c
	c := Cuadrado(10)
	t := Triangulo(5.5)

	// llamar al metodo con mi variable declarada
	fmt.Println(c.Area(), t.Area())
}
