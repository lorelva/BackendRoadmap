package main

import (
	"fmt"
)

func main() {
	port := 3000
	//implicit initiallization syntax
	isStarted := startWebServer(port)
	fmt.Println(isStarted)

	num1 := 0

	num1, n2, s1, b1, e1 := multi7(1, 2, "a")
	fmt.Println(num1, n2, s1, b1, e1)

	response, err := post(url)
	if err != nil {
		return nil, err
	}

	response, err = get(url)
	if err != nil {
		return nil, err
	}
}

// STEP 1: specify the data type of returning, in this case: bool
func startWebServer(port int) bool {
	fmt.Println("Starting server:")
	fmt.Println("Server started at port:", port)
	return true

}

// Syntaxis
// func {nombre de funcion}( {...n parametros {tipo de datos} } ) {n...valor de retorno} {
//	....
//}

func multi(num int) int {
	return num * num
}

func multi2(num int) (resultado int) {
	resultado = num * num
	return resultado
}

func multi3(num int) (int, int) {
	return num * num, num
}

func multi4(num int) (resultado int, numero_original int) {
	numero_original = num

	resultado = num * num
	return
}

func multi5(num int) (resultado int, numero_original int) {
	numero_original = num
	resultado = num * num
	return resultado, numero_original
}

func multi6(num int) (resultado, numero_original int) {
	numero_original = num
	resultado = num * num
	return resultado, numero_original
}

func multi7(num, num2 int, saludo string) (int, int, string, bool, error) {
	return 0, 1, "hola", false, nil
}
