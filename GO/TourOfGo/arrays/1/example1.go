package main

import "fmt"

func main() {
	//asign size
	var names [2]string
	//declare
	names[0] = "Christian"
	names[1] = "Lorena"

	//print the names separate(by the position of array)
	fmt.Println(names[0], names[1])
	//print the complete array
	fmt.Println(names)

	//other way to do this array
	names2 := [2]string{"Hernandez", "Valle"}
	fmt.Println(names2)

}
