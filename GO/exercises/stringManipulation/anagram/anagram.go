package main

import "fmt"

func main() {
	fmt.Println(isAnagram("silent", "listen"))

}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	charCount := [26]int{}
	for _, ch := range str1 {
		charCount[ch-'a']++
	}
	for _, ch := range str2 {
		charCount[ch-'a']--
		if charCount[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

/*
    Length Check:
        If the lengths of str1 and str2 are different, they cannot be anagrams, so return false.

    Character Count Array:
        A fixed-size array charCount of length 26 is used to track the frequency of each character in str1 and str2.
        Since the problem assumes lowercase English letters ('a' to 'z'), we can use the difference ch - 'a' to map each character to an index in the range 0 to 25.

    Increment Character Counts for str1:
        For each character in str1, increment the corresponding index in the charCount array.

    Decrement Character Counts for str2:
        For each character in str2, decrement the corresponding index in the charCount array.
        If the count becomes negative, it means str2 has more occurrences of that character than str1, so return false.

    Return True:
        If no mismatches are found and all counts are balanced, the strings are anagrams, and return true.

Example Walkthrough
Input: str1 = "listen", str2 = "silent"

    Step 1: Lengths are equal, so proceed.
    Step 2: Initialize charCount as an array of 26 zeros.
    Step 3: Count characters in str1:
        'l': charCount[11] becomes 1.
        'i': charCount[8] becomes 1.
        's': charCount[18] becomes 1.
        't': charCount[19] becomes 1.
        'e': charCount[4] becomes 1.
        'n': charCount[13] becomes 1.
    Step 4: Count characters in str2:
        's': charCount[18] becomes 0.
        'i': charCount[8] becomes 0.
        'l': charCount[11] becomes 0.
        'e': charCount[4] becomes 0.
        'n': charCount[13] becomes 0.
        't': charCount[19] becomes 0.
    Step 5: All counts are balanced, so return true.

*/
