package main

import (
	"fmt"
	"strings"
)

func main() {
	//FORMAT HHMMSS12h = "03:04:05PM"
	//TO 24h OutputString := "15:04:05"
	s := "03:04:05PM"
	conversion := timeConversion(s)
	fmt.Println(conversion)
}
func timeConversion(s string) string {
	separateTime := strings.Split(s, ":")
	//convert to 3 slices 03:04:05PM
	hours := separateTime[0]
	minutes := separateTime[1]
	seconds := separateTime[2]
	//separate AM OR PM from seconds[2]
	format := s[len(s)-2:]
	fmt.Println(hours)
	fmt.Println(minutes)
	fmt.Println(seconds)
	fmt.Println(format)

	var toHour int
	//validate if format  it's AM OR PM
	if format == "AM" && hours == "12" {
		hours = "00"
		//toHour, _ = strconv.Atoi(hours)

	}
	if format == "PM " && hours != "12" {
		hours += "12"
	}
	return s
}
