package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Println(result)

}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	duplicate := make(map[int]struct{})

	for _, number := range nums {
		if _, ok := duplicate[number]; ok {
			return true
		}

		duplicate[number] = struct{}{}
	}

	return false

}
