package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// swap function returns two strings and changes the position
func main() {
	word1, word2 := swap("hello", "world")
	fmt.Println(word1, word2)
}
