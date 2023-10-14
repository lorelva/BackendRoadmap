package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := "21"
	tonumber, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("There's an error")
	} else {
		fmt.Println(tonumber)
	}
}
