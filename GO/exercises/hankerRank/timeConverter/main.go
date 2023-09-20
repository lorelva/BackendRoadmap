package main

import "fmt"

func main() {

	s := "07:05:45PM"
	conversion := timeConversion(s)
	fmt.Println(conversion)

}

func timeConversion(s string) string {
	return s

}
