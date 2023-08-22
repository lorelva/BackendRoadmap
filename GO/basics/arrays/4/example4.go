package main

import "fmt"

func main() {
	//Calcualte the length of an array
	distros := []string{"Ubuntu", "Fedora", "ArchLinux", "Mint", "Debian"}
	fmt.Println(len(distros))

	//Change a value of a position in an array
	cars := [3]string{"Volvo", "Ford", "Mercedes"}
	cars[2] = "Ferrari"
	fmt.Println(cars)
}
