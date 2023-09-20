package main

import "fmt"

func main() {

	sum := 0

	//init statement--> i :=0
	//conditional expression--> i <10   these means numbers between 0 and less than 10
	//post statement -->executed at the end of every iteration

	//RULE: NO parentheses surrounding the for statement
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}
