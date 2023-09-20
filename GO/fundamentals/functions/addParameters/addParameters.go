package main

import "fmt"

func main() {
	fmt.Println("Testing a function")
	port := 3000
	//provide the variable to the function
	startWebServer(port)
}

// prepare the function to revieve data as parameter
func startWebServer(port int) {
	fmt.Println("Starting server:")
	//add a kind of functionallity with the variable created port
	fmt.Println("Server started at port:", port)
}
