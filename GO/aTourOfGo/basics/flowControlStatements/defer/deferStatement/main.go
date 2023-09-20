package main

import "fmt"

// DEFER is executed until the surrounding function returns.
// It means, it does everything first  and at last it will executed the content
// that the defer statement has
// It doesn't matter if it's declared at the beginning, it will executed at the ending
func main() {
	defer fmt.Println("METAL MUSIC")

	fmt.Println("I LOVE AND ENJOY")

}
