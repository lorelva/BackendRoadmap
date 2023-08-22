package main

import "fmt"

type Triangulito struct {
	Base   float64
	Altura float64
}

// Metodo para obtener el area con un struct como tipo de dato para mi metodo
func (t Triangulito) Area() float64 {
	return t.Base * t.Altura / 2
}

// Se agrega el * (puntero de memoria) que hace referencia a Triangulito para cambiar los valores
// Solo si se cambia un valor del struct, usar puntero
func (t *Triangulito) CambioDeValores(b, a float64) {
	t.Base = b
	t.Altura = a
}

func main() {
	t := Triangulito{
		Base:   4.5,
		Altura: 8.7,
	}

	fmt.Println(t.Altura, t.Base, t.Area())
	t.CambioDeValores(13.4, 20.1)
	fmt.Println(t.Altura, t.Base, t.Area())
}
