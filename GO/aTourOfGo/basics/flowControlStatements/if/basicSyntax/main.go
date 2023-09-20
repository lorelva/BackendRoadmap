package main

import "fmt"

func main() {

	//Greater than
	if 40 > 28 {
		fmt.Println("40 is greater than 28")
	}

	//Less than
	if 23 < 100 {
		fmt.Println("23 is less than 100")
	}

	//Not equal to
	x := 10
	y := 5

	name := "Lorena"
	name2 := "Lorelei"
	fmt.Println(x != y)

	//Equal to
	if x == y {
		fmt.Println("It's equal")
	}

	//Logical AND &&
	if x != y && name == name2 {

	}

	//Logical OR ||
	if name != name2 || x < y {
		fmt.Println("One condition it's true and the other not")
	} else {
		fmt.Println("Both are true")
	}

}
