package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//FORMAT HHMMSS12h = "03:04:05PM  TO 24h OutputString := "15:04:05"
	s := "06:56:45PM"
	conversion := timeConversion(s)
	fmt.Println(conversion)
}
func timeConversion(s string) string {
	separateTime := strings.SplitAfter(s, ":")
	//convert to 3 slices  the hour -->  06:56:45APM
	hours := separateTime[0]
	minutes := separateTime[1]
	secondsFormat := separateTime[2]
	//separate AM OR PM from seconds--> slice[2]
	seconds := secondsFormat[:2]
	format := secondsFormat[2:]
	//Conversion string to int for hours and minutes
	hoursConv, err := strconv.Atoi(hours[:2])
	if err != nil {
		fmt.Println("There is an error:", err)
		return ""
	}
	//validate if format  it's AM OR PM
	if format == "AM" && hoursConv == 12 {
		hoursConv = 0
	} else if format == "PM" && hoursConv != 12 {
		hoursConv += 12
	}
	/*hoursReturn := strconv.Itoa(hoursConv)
	minutesReturn := strconv.Itoa(minutesConv)
	*/
	hours = fmt.Sprintf("%02d:", hoursConv)
	//hours := strconv.Itoa(hoursConv)
	return hours + minutes + seconds
}
