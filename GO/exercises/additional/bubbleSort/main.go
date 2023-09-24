/*
Problema : Ordenación por Burbuja

El algoritmo de ordenación por burbuja funciona comparando pares de elementos adyacentes
en una lista y los intercambia si están en el orden incorrecto.
Se repite este proceso hasta que no se requieran más intercambios, lo que significa que la
lista está ordenada.

Instrucciones:

    - Comienza en el primer elemento de la lista.
    - Compara el elemento actual con el siguiente.
    - Si están en el orden incorrecto, intercámbialos.
    - Avanza al siguiente par de elementos y repite los pasos 2-3.
    - Continúa hasta que llegues al final de la lista.
    - Repite los pasos 1-5 hasta que no se requieran más intercambios.
*/

/*
	PASOS PARA RECORRER Y ORDENAR:
	1.-Recorrer el arreglo desordenado  for len(arr)
	2.-Saber las dos primeras posiciones para hacer la comparación de si pos[0] > pos[1]
	3.-Si pos[0] > pos[1] --> intercambiar de posicion el valor dentro de estas dos posiciones
	4.- Si no , dejar mismos valores y tomar en cuenta la posición pos[2]
	5.- Si pos[1](posicion anterior)> pos[2](actual) intercambiar los valores de las posiciones

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//arreglo := []int{5, 3, 6, 1, 2}
	arreglo := generateRandomValues(10)
	//ordenados: [1, 2 , 3, 5, 6]
	fmt.Printf("Arreglo sin ordernar: %v \n", arreglo)
	// Mandar el arreglo a la función y una vez ordenado se debe imprimir
	arreglo2 := ordenarArreglo(arreglo)
	fmt.Println(arreglo2)

}

// Crear una función que retorne el arreglo ordenado y debe recibir como parametro el arreglo desordenado
// arreglo2 es la array que le llega por parametro(el desordenado)
// retornar un arreglo también pero este va a ser con la finalidad de ya tenerlos ordenados
func ordenarArreglo(arreglo []int) []int {
	fmt.Println("LOngitud: ", len(arreglo))

	//se usa el valor del parametro: arreglo2 , para recorrerlo
	for i := 0; i < len(arreglo); i++ {
		for k := 0; k < len(arreglo)-1; k++ {
			//fmt.Println(arreglo)
			if arreglo[k] > arreglo[k+1] {
				actual := arreglo[k]
				siguiente := arreglo[k+1]
				//{1, 2, 3, 5, 6}
				// 5>3 -->true
				arreglo[k] = siguiente
				arreglo[k+1] = actual

			}
		}

	}
	return arreglo
}

func generateRandomValues(n int) []int {
	rand.Seed(time.Now().UnixNano())
	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = rand.Intn(10001) // Genera valores entre 0 y 10,000
	}
	return values
}
