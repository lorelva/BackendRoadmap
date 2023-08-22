package main

import "fmt"

func main() {

	//declarar salon
	salon := make(map[string]int)

	//asignar valores
	salon["3A"] = 23
	salon["7B"] = 30
	salon["4C"] = 27
	salon["6A"] = 20
	salon["1D"] = 42
	salon["5B"] = 31

	//verificar
	num, ok := salon["4C"]
	if ok {
		fmt.Println("Numero de alumnos en el salon  es:")
		fmt.Println(num)
	} else if !ok {
		fmt.Println("Salon no encontrado en la lista")

	}

	_, ok = salon["7B"]
	if ok {
		fmt.Println("Salon no encontrado en la lista")
	}
}
