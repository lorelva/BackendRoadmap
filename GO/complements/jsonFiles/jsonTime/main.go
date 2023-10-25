package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Time struct {
	TimeValue string `json:"time_value,omitempty"`
}

func main() {

	content, err := os.ReadFile("./time.json")
	if err != nil {
		log.Fatal("There was an error", err)
	}

	//t := Time{}
	t := Time{}
	json.Unmarshal(content, &t)

	err = json.Unmarshal(content, &t)

	if err != nil {
		log.Fatal("Error unexpected:", err)
	}

	convertedTime := timeConversion(t)
	fmt.Printf("Converted Time: %s\n", convertedTime)
}

func timeConversion(t Time) string {

	separateTime := strings.SplitAfter(t.TimeValue, ":")
	hours := separateTime[0]
	minutes := separateTime[1]
	secondsFormat := separateTime[2]
	seconds := secondsFormat[:2]
	format := secondsFormat[2:]

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

	hours = fmt.Sprintf("%02d:", hoursConv)
	return hours + minutes + seconds
}
