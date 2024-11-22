package main

import "fmt"

// STRING INDEXING Direct access via str[index].
// Specific characters.

func wordIndexing() {
	word := "SYMPHONIC"
	word2 := "METAL"

	fmt.Println("Fisrt character:", string(word[0]))
	fmt.Println("Character of the index 2:", string(word[2]))
	fmt.Println("Last character:", string(word[len(word)-1]))

	fmt.Println("\nAccess to first character of word2: ", string(word2[0]))
	fmt.Println("Character of the index 3: ", string(word2[3]))
	fmt.Println("Last character: ", string(word2[len(word2)-1]))

	completeWord := word + " " + word2
	fmt.Println(completeWord)

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
