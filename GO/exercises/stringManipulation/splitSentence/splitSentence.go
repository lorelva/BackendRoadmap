package main

import (
	"fmt"
	"strings"
)

func splitSentence(sentence string) []string {
	words := strings.Fields(sentence)
	return words
}

func main() {
	sentence := "This is a test sentence."
	words := splitSentence(sentence)

	for _, word := range words {
		fmt.Println(word)
	}
}
