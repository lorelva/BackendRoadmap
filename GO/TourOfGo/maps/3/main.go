package main

import "fmt"

func main() {

	//declarar mapa para las edades
	edades := make(map[string]int)

	//asignar valores
	edades["Chrsitian"] = 21
	edades["Karla"] = 26
	edades["Lorena"] = 21
	edades["David"] = 18
	edades["Juan"] = 25

	//verificar si está un valor
	nombre := "Christian"
	v, ok := edades[nombre]
	if ok {
		fmt.Printf("Edad de %s es %d \n", nombre, v)
	} else {
		fmt.Printf("%s no se encuentra en la lista.\n", nombre)
	}

	//verificar si no  está un valor
	nombre2 := "Daniel"
	v2, ok := edades[nombre2]
	if ok {
		fmt.Printf("La edad de %s es %d años.\n", nombre2, v2)
	} else {
		fmt.Printf("%s no se encuentra en la lista.\n", nombre2)
	}

}
