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
	//cancionBuscar := "Jaded"
	//cancionBuscar := "Those nights"
	//resultado := buscarCancion(listaCancion, cancionBuscar)
	//resultado := fmt.Println(buscarCancion(listaCancion, "Those nights"))
	//fmt.Printf("La canción introducida %v se encontró: %v\n", cancionBuscar, resultado)
	fmt.Println(buscarCancion(listaCancion, "Those nights"))

}

// función para buscar el elemento deseado
func buscarCancion(listaCancion []string, cancionBuscar string) bool {
	/*for i := 0; i < len(listaCancion); i++ {
		if cancionBuscar == listaCancion[i] {
			return true
		}
	}
	*/
	for _, v := range listaCancion {
		if v == cancionBuscar {
			return true
		}
		//fmt.Println(listaCancion[i])
	}
	return false
}
