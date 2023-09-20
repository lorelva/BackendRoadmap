/*

**Explicación de Búsqueda Binaria:** Imagina que estás buscando una palabra en un diccionario.
En lugar de comenzar desde la primera página y revisar una por una hasta encontrar la palabra,
abres el diccionario por la mitad. Si la palabra que buscas está antes (en orden alfabético) de
la página donde abriste, sabes que tienes que buscar en la primera mitad del diccionario. Si está
después, buscas en la segunda mitad. Luego, tomas la mitad en la que decidiste buscar y la divides
nuevamente, repitiendo el proceso. De esta manera, reduces rápidamente la cantidad de páginas que
necesitas revisar. Eso es básicamente cómo funciona la búsqueda binaria, pero con números en una
lista ordenada.
**Instrucciones sencillas para Búsqueda Binaria:**
 1. Determina el punto posicion de
la lista. 2. Mira el valor en ese punto posicion. 3. Si el valor es el que estás buscando, ¡ya lo
encontraste! 4. Si el valor es menor que lo que buscas, repite el proceso pero solo con la mitad
de la derecha (los números mayores). 5. Si el valor es mayor que lo que buscas, repite el proceso
pero solo con la mitad de la izquierda (los números menores). 6. Si ya no quedan más números en
la lista y no encontraste lo que buscabas, significa que el número no está en la lista. Siguiendo
estos pasos, puedes encontrar rápidamente un número en una lista ordenada sin tener que revisar
cada número uno por uno. ¡Es mucho más rápido!
*/

package main

import "fmt"

func main() {
	array := []int{1, 4, 5, 7, 9, 12, 15, 18, 19, 22, 25, 29, 36, 41}
	numeroEncontrar := 22

	result := busquedaBinaria(array, numeroEncontrar)
	fmt.Println(result)
}

func busquedaBinaria(array []int, numeroEncontrar int) int {
	//para declarar las posiciones
	inicio := 0
	final := len(array) - 1
	//posicion := (inicio + final) / 2
	//fmt.Println(posicion)

	for i := 0; i <= len(array)-1; i++ {
		posicion := (inicio + final) / 2
		numeroEncontrar := array[posicion]

		if posicion == numeroEncontrar {
			fmt.Printf("Número %d encontrado", numeroEncontrar)
		}
		if posicion < numeroEncontrar {
			inicio = posicion + 1
		}
		if posicion > numeroEncontrar {
			final = posicion - 1
		}
		if posicion != numeroEncontrar {
			fmt.Printf("Numero %d no encontrado", numeroEncontrar)
		}
	}
	return numeroEncontrar

}
