package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CONVERT JSON FILE TO GO STRUCT
// declare go struct
type CustomTime struct {
	TimeValue string `json:"time_value,omitempty"`
}

func main() {
	//READ JSON FILE
	content, err := os.ReadFile("./time.json")
	if err != nil {
		log.Fatal(err)
	}

	//show the content of the json file
	//fmt.Printf("File contents: %s", content)

	t := CustomTime{}
	//json.Unmarshal(content, &t)

	err = json.Unmarshal(content, &t)
	if err != nil {
		log.Fatal("Error JSON Unmarshalling:", err)
	}

	fmt.Println(t)

	//CONVERT STRUCT TO JSON
	convertedTime := timeConversion(t)
	resultado := CustomTime{
		TimeValue: convertedTime,
	}

	resultEncoded, err := json.Marshal(resultado)
	if err != nil {
		log.Fatalf("Error: %v obtained", err)

	}

	fmt.Printf("Converted Time: %s\n", string(resultEncoded))
}

func timeConversion(t CustomTime) string {

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
