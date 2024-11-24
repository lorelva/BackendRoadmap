package main

import "fmt"

func main() {
	nums := make([]int, 0, 0)
	for i := 1; i <= 5; i++ {
		nums = append(nums, i)
		fmt.Println(nums)
	}

	music()

}

func music() {
	metal := make(map[string]string)
	metal["death"] = "Deicide"
	metal["trash"] = "Kreator"
	metal["heavy"] = "Pantera"
	metal["black"] = "Astarte"
	metal["symphonic"] = "Epica"

	fmt.Printf(metal["trash"])
}
