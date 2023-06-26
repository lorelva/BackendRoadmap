package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Number %v it's double is  %v\n", v, v+v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("Is it true or false ? %T!\n", v)
	}
}

func main() {
	do(10)
	do("hello")
	do(true)
}
