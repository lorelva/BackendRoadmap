package main

import "fmt"

func main() {

	var resultado interface{}
	fmt.Println("INgrese el valor:")
	fmt.Scan(&resultado)
	fmt.Printf("Tipo de datos: %T el valor es: %v\n", resultado, resultado)

}
