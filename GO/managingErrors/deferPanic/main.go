package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		//recover se utiliza para capturar un panic y manejarlo de manera controlada
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

//DEFER: retrasar la ejecución de una función hasta que la función circundante retorne
//PANIC: detiene el flujo ordinario de control y comienza a entrar en pánico(mensaje)
//RECOVER :se utiliza para capturar un panic y manejarlo de manera controlada
