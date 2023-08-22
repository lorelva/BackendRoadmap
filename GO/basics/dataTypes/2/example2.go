package main

import "fmt"

// Variables with initializers

var i, j int = 1, 2 //this way or separate
//var i int = 1
//var j int = 2

func main() {
	var c, python, java = true, false, "NO!"
	fmt.Println(i, j, c, python, java)
}
