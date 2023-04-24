package main

import "fmt"

func main() {
	port := 3000
	//implicit initiallization syntax
	isStarted := startWebServer(port)
	fmt.Println(isStarted)
}

// STEP 1: specify the data type of returning, in this case: bool
func startWebServer(port int) bool {
	fmt.Println("Starting server:")
	fmt.Println("Server started at port:", port)
	return true
}
