package main

import "fmt"

// STRING INDEXING Direct access via str[index].
// Specific characters.

func wordIndexing() {
	word := "LOVELY" + "GLAMOUR"

	fmt.Println(string(word[0]))
	fmt.Println(string(word[1]))
	fmt.Println(string(word[11]))
	fmt.Println(string(word[3]))
	fmt.Println(string(word[6]))
	fmt.Println(string(word[10]))

}

func main() {
	wordIndexing()
	/*numbers := []int{1, 56, 4, 37, 90, 13}
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 20 {
			fmt.Println(i, "IT PASSED")
		} else {
			fmt.Println(i)
		}
	}
	*/
}
