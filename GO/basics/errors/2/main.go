package main

import (
	"errors"
	"fmt"
)

func encuentraFrase(s string) error {
	sl := []string{"hola", "adios"}

	for _, v := range sl {
		if s == v {
			return nil
		}
	}

	return errors.New("error, no existe la frase")
}

func main() {
	err := encuentraFrase("hola")
	if err != nil {
		fmt.Println(err)
	}
}
