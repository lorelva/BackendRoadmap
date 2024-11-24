package main

import "fmt"

func main() {

	fmt.Println(areIsomorphic("green", "kreep"))

}

func areIsomorphic(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	mappingStr1 := [256]int{}
	mappingStr2 := [256]int{}

	for i := range str1 {
		if mappingStr1[str1[i]] != mappingStr2[str2[i]] {
			return false
		}
		mappingStr1[str1[i]] = i + 1
		mappingStr2[str2[i]] = i + 1
	}
	return true
}

/*
How It Works

    Check Lengths:
        If str1 and str2 have different lengths, they can't be isomorphic, so it immediately returns false.

    Character Mapping:
        mappingStr1 and mappingStr2 are arrays of size 256 (for ASCII characters). These arrays track the last-seen index for each character in str1 and str2.
        For each character pair (str1[i], str2[i]), if the last-seen indices from both mappings don't match, the strings aren't isomorphic.

    Update Mapping:
        The indices in mappingStr1[str1[i]] and mappingStr2[str2[i]] are updated with i + 1 to avoid conflicts with the default 0 value.

    Return True:
        If all character pairs can be mapped consistently, the function returns true.

Example Walkthrough
Input: str1 = "egg", str2 = "add"

    Step 1: Lengths are equal, so proceed.
    Step 2: Initialize mappingStr1 and mappingStr2 as arrays of 256 zeros.
    Step 3: Iterate through each index i:
        i = 0: str1[0] = 'e', str2[0] = 'a'
            mappingStr1['e'] == 0 matches mappingStr2['a'] == 0.
            Update: mappingStr1['e'] = 1, mappingStr2['a'] = 1.
        i = 1: str1[1] = 'g', str2[1] = 'd'
            mappingStr1['g'] == 0 matches mappingStr2['d'] == 0.
            Update: mappingStr1['g'] = 2, mappingStr2['d'] = 2.
        i = 2: str1[2] = 'g', str2[2] = 'd'
            mappingStr1['g'] == 2 matches mappingStr2['d'] == 2.
            Update: mappingStr1['g'] = 3, mappingStr2['d'] = 3.
    Step 4: All checks pass, so return true.

*/
