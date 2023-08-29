package main

import "fmt"

// hacer una busqueda lineal del numero en el arreglo,
func main() {
	//declarar los parametros a utilizar para buscar un numero dentro de un arreglo
	listaNumero := []int{6, 7, 20, 89, 34, 51}
	// si existe
	numDestino := 18

	resultado := buscarNumero(listaNumero, numDestino)
	fmt.Printf("El numero: %d se encontró: %v\n", numDestino, resultado)

}

// función para buscar el elemento deseados
// recibe como parametro el array de los numeros, el numero a encontrar y va a devolver con un true or false si se encuentra
func buscarNumero(listaNumero []int, numDestino int) bool {
	for i := 0; i < len(listaNumero); i++ {
		for k := 0; k < len(listaNumero); k++ {
			if numDestino == listaNumero[k] {
				return true
			}
		}
		//fmt.Println(listaNumero[i])
	}
	return false
}
