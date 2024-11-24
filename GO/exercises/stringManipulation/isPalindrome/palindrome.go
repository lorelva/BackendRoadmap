package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	// Step 1: Normalize the string
	newString := ""
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			newString += string(unicode.ToLower(char))
		}
	}

	// Step 2: Check if the newString string is a palindrome
	for i := 0; i < len(newString)/2; i++ {
		if newString[i] != newString[len(newString)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	// Example test cases
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
}
