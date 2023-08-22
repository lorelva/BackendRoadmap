package main

import "fmt"

// Iterating over an array using range
func main() {

	//days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	var days [7]string
	days[0] = "Sunday"
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"
	days[6] = "Saturday"

	for index, value := range days {
		fmt.Printf("Day %d of week = %s\n", index, value)

	}
}
