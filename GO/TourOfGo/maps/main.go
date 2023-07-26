package main

import "fmt"

func main() {

	//map[KeyType]ValueType
	// nunca se debe crear un mapa vacio, siempre usar var, make o si se usa := se debe poner vacio o con un valor

	mapasV1()
	mapasV2()

}

func mapasV1() {
	//declarar e inicializar mapa

	// nunca se ocupa esta forma porque se crea vacio (nil)
	// var diccionario map[string]string

	// forma correcta con make
	diccionario := make(map[string]string)

	// asignar un valor al mapa
	diccionario["lapiz"] = "Un lapiz es un coso de un árbol que contiene grafito"

	// acceder solo al valor
	definicion := diccionario["lapiz"]
	fmt.Println(definicion)

	definicion2, ok := diccionario["libro"]
	if ok {
		fmt.Println("La llave loibro existe")
		fmt.Println(definicion2)
	} else if !ok {
		fmt.Println("La llave no existe")

	}

	_, ok = diccionario["color"]
	if !ok {
		fmt.Println("La llave color no es¿xiste")
	}

}

//edades de personas
//valores de divisas mnx a dolar
//numero de niños o morros en un salon

func mapasV2() {

}
