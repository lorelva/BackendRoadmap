package main

import "fmt"

// Short variable declarations
func main() {
	var i, j int = 1, 2
	k := 3 //Declaration with implicit type.
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
