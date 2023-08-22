package main

//import new library
import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) error {
	fmt.Println("Starting server:")
	fmt.Println("Server started at port:", port)
	return errors.New("Something went wrong")
}
