package main

import "fmt"

func main() {

	input := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	//input := []int64{1, 3, 5, 2}
	//sumar el valor total de los elementos que tiene el array
	result := aVeryBigSum(input)
	fmt.Println("Resultado de la suma:", result)
}

func aVeryBigSum(input []int64) int64 {
	// Iniciliazar la suma desde cero
	var total int64 = 0

	//hacer un for range para los elementos del array
	for _, element := range input {
		//EJEMPLO  total = 0
		// total + element(1)= 1
		//total = 1+3 = 4
		//total = 4+5 = 9
		//total = 9+ 2 = 11
		//debe dar como resultado 11
		total = total + element
	}
	return total

}
