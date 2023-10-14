package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, This is my, first code in, Golang!"
	//Split string without removing separator(comas)
	split := strings.SplitAfter(str, ",")
	fmt.Println(split)

}
