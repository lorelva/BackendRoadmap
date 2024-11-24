package main

import "fmt"

func main() {

	fmt.Println(isPangram("Sphinx of black quartz, judge my vow"))

}

func isPangram(str string) bool {
	letterArray := [26]bool{}
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			letterArray[ch-'a'] = true
		} else if ch >= 'A' && ch <= 'Z' {
			letterArray[ch-'A'] = true
		}
	}
	for _, exists := range letterArray {
		if !exists {
			return false
		}
	}
	return true
}
