package main

import (
	"fmt"
	"runtime"
)

// only runs the selected case, not all the cases that follow.
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("Windows")
		fmt.Printf("%s.\n", os)
	}
}
