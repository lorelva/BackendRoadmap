package main

import (
	"fmt"
	"strings"
)

func main() {

	timeToConvert := "08:05:56PM"

	result := strings.Split(timeToConvert, ":")
	fmt.Println(result)

}
