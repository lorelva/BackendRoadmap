package main

import "fmt"

//Type inference

func main() {

	var i int
	j := i // j is an int

	//type int
	v := 42

	//type float
	f := 23.67

	//type complex
	c := 0.867 + 0.5i

	//type bool
	b := true

	fmt.Printf("Type %T\n", v)
	fmt.Printf("Type %T\n", f)
	fmt.Printf("Type %T\n", c)
	fmt.Printf("Type %T\n", b)

	fmt.Printf("Type %T\n", j)
}
