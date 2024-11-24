package main

import (
	"fmt"
)

func isIsogram(str string) bool {
	if len(str) == 0 { // Step 1: Handle empty strings
		return false
	}

	charArray := [26]bool{} // Step 2: Array for tracking letters

	for _, ch := range str { // Step 3: Loop through the string
		index := ch - 'a'     // Calculate the index
		if charArray[index] { // Check for duplicates
			return false
		}
		charArray[index] = true
	}
	return true // Step 4: Return true if no duplicates
}

func main() {
	fmt.Println(isIsogram("machine"))
	fmt.Println(isIsogram("kreator"))
	fmt.Println(isIsogram("abcdefg"))
}

/*

The string s is iterated through, character by character.
For each character (char), it checks whether it is a letter or a digit using unicode.IsLetter(char) or unicode.IsDigit(char).
If the character is alphanumeric, it is converted to lowercase using unicode.ToLower(char) and appended to the new string newString.
After this loop, newString contains only lowercase letters and digits, with all other characters removed.

*/

/*
Goal: The goal of this block is to check if the newString is a palindrome.
Process:

    The loop runs from i = 0 to i < len(newString)/2. This means it only checks the first half of the string because if the string is a palindrome, the second half should mirror the first half.
    Inside the loop, it compares the i-th character from the beginning (newString[i]) with the i-th character from the end (newString[len(newString)-1-i]).
    If any pair of characters do not match, the function immediately returns false, indicating the string is not a palindrome.

*/
