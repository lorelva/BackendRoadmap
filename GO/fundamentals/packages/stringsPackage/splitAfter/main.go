package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	secondExample()
}

func secondExample() {
	str := "Hello, This is my, first code in, Golang!"
	//Split a string by separator, this case the coma
	split := strings.Split(str, ",")
	fmt.Println(split)
}
