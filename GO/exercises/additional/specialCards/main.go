package main

import "fmt"

func main() {
	k := []int{2, 3, 1, 3}
	//dont sum the numbers that are repeated so...
	//special cards 2 , 1 --> 2+1 = 3
	result := solve(k)
	fmt.Println(result)
}

func solve(k []int) int {
	//para buscar  cada numero repetido en un mapa
	find := make(map[int]int)
	//para sumar los numeros no repetidos
	sum := 0
	//iterar los numeros y si se encuentra repetido y en value se van guardando
	for _, value := range k {
		find[value]++
		//fmt.Println(value)
		//sum += k[value]
	}
	for _, value := range k {
		//1 para saber que es un valor unico y se deberá sumar
		if find[value] == 1 {
			//hacer la suma de los valores no repetidos y guardarlos en sum
			sum += value
		}
	}
	return sum
}
