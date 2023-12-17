package main

import "fmt"

// Given a string , return it reverse...
// for example:
// cat --> tac
// hello --> olleh
// lorena --> anerol

func main() {

	word := "lorena"
	toReverse := makeReverse(word)
	fmt.Println(toReverse)

}

func makeReverse(word string) string {

	reverse := ""
	for _, value := range word {

		// Concatenate the current character at the beginning of the reversed string.
		//Concatena el carácter actual al inicio de la cadena invertida
		reverse = string(value) + reverse

	}

	return reverse

}
