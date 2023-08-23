package main

import "fmt"

// VARIBLAES WITH VAR
// var statement declares a list of variables, and the type of data is last
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
	variables()
}

// VARIABLES WITH INITITIALIZERS
// var declaration can include initializers, one per variable:
var j, k int = 1, 2

func variables() {
	//declare new values with var decalration and initializers
	//It's not necessary to declare the data type it would take it for default on what you initialize
	var c, python, java = 23, "Programming language for AI", true
	fmt.Println(j, k, c, python, java)
}
