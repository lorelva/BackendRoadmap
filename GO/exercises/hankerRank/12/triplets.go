package main

import "fmt"

func main() {

	a := []int32{1, 2, 3}
	b := []int32{3, 2, 1}

	total := compareTriplets(a, b)
	fmt.Println(total)

}

func compareTriplets(a []int32, b []int32) []int32 {

	var alicePoints int32 = 0
	var bobPoints int32 = 0

	for i := 0; i < len(a); i++ {
		/*for j := 0; j < len(b); j++ {*/

		if a[i] > b[i] {
			alicePoints++
		}
		if a[i] < b[i] {
			bobPoints++
		}
	}
	return []int32{alicePoints, bobPoints}
}
