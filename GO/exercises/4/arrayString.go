package main

import "fmt"

// hacer una busqueda lineal de una cadena de palabras en el arreglo,
func main() {
	//declarar los parametros a utilizar para buscar el nombre de una cancion dentro de un arreglo
	listaCancion := []string{
		"I get",
		"Comatose",
		"Those Nights",
		"Whispers",
		"Last night",
		"Yours to",
		"Falling ",
		"Say Goodbye",
		"Karma",
	}
	// si existe
	cancionBuscar := "Jaded"
	/*var cancionBuscar2 string
	fmt.Println("Introduce el nombre de una canción: ")
	fmt.Scanln(&cancionBuscar2)
	*/
	resultado := buscarCancion(listaCancion, cancionBuscar)
	fmt.Printf("La canción introducida %v se encontró: %v\n", cancionBuscar, resultado)

}

// función para buscar el elemento deseado
func buscarCancion(listaCancion []string, cancionBuscar string) bool {
	for i := 0; i < len(listaCancion); i++ {
		for k := 0; k < len(listaCancion); k++ {
			if cancionBuscar == listaCancion[k] {
				return true
			}
		}
		//fmt.Println(listaCancion[i])
	}
	return false
}
