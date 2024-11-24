package main

import (
	"fmt"
	"strings"
)

func main() {

	str := ".lyah_valle_1268_gmail.com"

	newString := strings.Replace(str, "_", "@", 1)
	fmt.Println("New String : ", newString)

	newString2 := strings.ReplaceAll(str, "_", " ")
	fmt.Println("New String : ", newString2)

	newString3 := strings.Replace(str, ".", "_", 2)
	fmt.Println("New String: ", newString3)

	splitString := strings.Split(str, "12")
	fmt.Println(splitString)
}
