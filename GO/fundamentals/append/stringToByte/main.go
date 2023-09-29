package main

import "fmt"

func main() {

	message := append([]byte("Hello "), "world!"...)
	fmt.Println(message)

}
