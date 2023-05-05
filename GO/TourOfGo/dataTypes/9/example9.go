package main

import "fmt"

// constants are decalred like variables but using the keyword: const
const Pi = 3.14
const MNX = 18.78

func main() {

	const World = "Merhaba"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	//extra example
	var convert float64 = 80
	var total = convert * MNX
	fmt.Println("TOTAL", total)
}
