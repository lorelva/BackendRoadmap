package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//FORMAT HHMMSS12h = "03:04:05PM  TO 24h OutputString := "15:04:05"
	s := "12:10:05AM"
	conversion := timeConversion(s)
	fmt.Println(conversion)
}

func timeConversion(s string) string {
	//convert to 3 slices 03:04:05PM
	separateTime := strings.SplitAfter(s, ":")
	hours := separateTime[0]
	minutes := separateTime[1]
	secondsFormat := separateTime[2]
	//separate AM OR PM from seconds[2]
	seconds := secondsFormat[:2]
	format := secondsFormat[2:]
	/*fmt.Println(hours)
	fmt.Println(minutes)
	fmt.Println(secondsFormat)
	fmt.Println(seconds)
	fmt.Println(format)
	*/

	//Conversion string to int for hours and minutes
	hoursConv, err := strconv.Atoi(hours[:2])
	if err != nil {
		fmt.Println("There is an error:", err)
		return ""
	}
	/*minutesConv, err := strconv.Atoi(minutes[:2])
	if err != nil {
		fmt.Println("There's an error:", err)
		return """
	}
	*/
	//fmt.Println(hoursConv)
	//fmt.Println(minutesConv)
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
	//minutes = fmt.Sprintf("%02d", minutesConv)
	return hours + minutes + seconds
}
