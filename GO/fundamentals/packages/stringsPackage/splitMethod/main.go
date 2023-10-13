package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "lorena-valle"
	//split
	//arguments--> variable , "what element to separate between the string"
	fmt.Println(strings.Split(name, "-"))

	phrase := "Everybody wants to rule , the world, the whole world"

	wholeSplit := strings.Split(phrase, ",")

	for word := range wholeSplit {
		fmt.Println(wholeSplit[word])

	}

}
