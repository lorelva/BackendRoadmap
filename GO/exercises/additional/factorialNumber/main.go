/*
Instrucciones sencillas usando loops:

	    Empieza con un número que quieras calcular, digamos, el número 5.
	    Multiplica ese número por el número anterior a él (4), y luego por el anterior a ese (3),
		y así sucesivamente, hasta llegar a 1.
	    Todos estos números multiplicados entre sí te dan el resultado. Para nuestro ejemplo, sería
		5 x 4 x 3 x 2 x 1 = 120.

Paso a Paso:

	Toma el número (por ejemplo, 5).
	Usa un loop que vaya desde 1 hasta ese número.
	En cada paso, multiplica el resultado actual por el número en ese paso.
	Cuando el loop termine, tendrás el factorial del número.
*/
package main

import "fmt"

func main() {
	//declarar el numero a buscar
	numero := 7
	// asignar el valor devuelto por la función  convertir Factorial con su argumento
	resultadoFinal := convertirFactorial(numero)
	fmt.Printf("El factorial del numero %d es: %v", numero, resultadoFinal)

}

func convertirFactorial(numero int) int {
	//Guardar el resultado del numero factorial
	resultado := 1
	//for para recorrer desde el numero 1 hasta el numero introducido
	for i := 1; i <= numero; i++ {
		//se tiene que acumular el producto(multiplicación) de los números
		resultado *= i
		//fmt.Println(resultado)
	}
	return resultado
}
