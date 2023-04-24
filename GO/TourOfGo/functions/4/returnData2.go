package main

import "fmt"

func main() {
	port := 3000
	//implicit initiallization syntax
	err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) error {
	fmt.Println("Starting server:")
	fmt.Println("Server started at port:", port)
	return nil
}
