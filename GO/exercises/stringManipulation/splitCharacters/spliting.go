package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Lorenao!"
	fmt.Println(strings.Split(word, "o"))
	fmt.Println(strings.Split(word, "r"))
}
