package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "Hello! How are! you?"
	re := regexp.MustCompile(`[!?.]`)
	cleaned := re.ReplaceAllString(str, "")
	fmt.Println("String without symbols:", cleaned)

	newString := strings.Split(str, "!")
	fmt.Println("String without symbols: ", newString)

}
