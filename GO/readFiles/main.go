package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	test1 := "2024-04-22"
	test2 := "2024-05-15"
	test3 := "2024-05-25"
	test4 := "2024-06-01"
	test5 := "2024-06-09"
	test6 := "2024-07-12"

	tests := []string{test1, test2, test3, test4, test5, test6}

	//read user date
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce a date with format YYYY-MM-DD:")

	date, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	date = strings.TrimSpace(date)

	testFile := ""

	for _, testDate := range tests {
		if testDate == date {
			testFile = date + "_databasecopy.txt"
			break
		}
	}

	if testFile == "" {
		log.Fatalf("Provided date %s not founded", date)
	}

	//read file
	data, err := os.ReadFile(testFile)
	if err != nil {
		log.Fatalf("Could not read the file %s: %v", testFile, err)
	}

	//show content of specific file
	fmt.Println(string(data))
}
