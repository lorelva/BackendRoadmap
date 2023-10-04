package main

import (
	"fmt"
)

func main() {

	s := "07:05:45PM"
	conversion := timeConversion(s)
	fmt.Println(conversion)

}

func timeConversion(s string) string {

	//FORMAT not defined in the time package : HHMMSS12h = "3:04:05 PM"
	stringtoConvert := s
	outputString := "19:05:45"

	return s

}
