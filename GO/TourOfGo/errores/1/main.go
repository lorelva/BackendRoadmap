package main

import (
	"fmt"
	"strconv"
)

func main() {

	for {
		var s string
		fmt.Println("ingresa un valor")
		fmt.Scanf("%s", &s)

		numero, err := strconv.Atoi(s)
		if err == nil {
			fmt.Printf("El numero es %v, y sin errores %v \n", numero, err)
		}
		if err != nil {
			fmt.Printf("No se puede convertir el valor %v a entero, el error es: %v\n", s, err)
		}
	}

}
