package main

import "fmt"

// Short variable declarations

func main() {
	var i, j int = 1, 2
	// ":= " statement: Can be only used inside a function,
	//it can be replaced instead of  var declaration with implicit type.
	k := 3
	//var c, python, java = true, false, "no!" --> Using var statement
	c, python, java := true, false, "no!" //--> No var statement --> replace for  := statement
	fmt.Println(i, j, k, c, python, java)
}
