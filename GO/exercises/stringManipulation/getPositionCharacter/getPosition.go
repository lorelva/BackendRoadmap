package main

import (
	"fmt"
	"strings"
)

func main() {

	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
	// ! I _ ! L o v e _ M u  s  _  i  c  !

	str2 := "!I_!Love_Mus_ic"
	char := '_'
	fmt.Println("First occurrence of '!':", strings.Index(str2, "!"))
	fmt.Println("Last occurrence of '!':", strings.LastIndex(str2, "!"))

	fmt.Printf("\nAll occurrences of '%c' in '%s':\n", char, str2)

	for index, position := range str2 {
		if position == char {
			fmt.Println("Position founded: ", index)

		}
	}
}
