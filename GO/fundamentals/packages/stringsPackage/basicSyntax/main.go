package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "lorena valle"
	//using package strings for the follwing methods:
	//UpperCase
	fmt.Println(strings.ToUpper(name))

	//TitleCase -->I have discovered that it deprecated ...
	fmt.Println(strings.Title(name))

}
