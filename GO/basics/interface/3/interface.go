package main

import "fmt"

type Cuadrado float64

type Triangulo struct {
	Base   float64
	Altura float64
}

func valoresDinamicos(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("Tipo %T, valor %v \n", i, i)

	case int:
		fmt.Printf("Tipo %T, valor %v \n", i, i)

	case float32:
		fmt.Printf("Tipo %T, valor %v \n", i, i)

	case bool:
		fmt.Printf("Tipo %T, valor %v \n", i, i)
	case Cuadrado:
		fmt.Println("cuadrado")
	default:
		fmt.Printf("Tipo desconocido %T, valor %v \n", i, i)

	}

}

func retornar(i interface{}) interface{} {
	return i
}

func main() {

	c := Cuadrado(10.2)
	t := Triangulo{
		Base:   12.78,
		Altura: 18.89,
	}

	valoresDinamicos("soy un string")
	valoresDinamicos(15.12)
	valoresDinamicos(23)
	valoresDinamicos(true)

	valoresDinamicos(c)
	valoresDinamicos(t)

	fmt.Println(retornar(10))
	fmt.Println(retornar(true))
	fmt.Println(retornar(10 * 2))

}
