package main

import "fmt"

func main() {

	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Moka", "Capuccino"},
		"tea":    {"Hot tea", "Chai Tea", "Chai Latte"},
	}

	fmt.Println(m["coffee"])
	m["other"] = []string{"Hot Chocolate"}
	fmt.Println(m)

	delete(m, "tea")
	fmt.Println(m)

	fmt.Println(m["tea"])
	v, ok := m["tea"]
	fmt.Println(v, ok)

	m2 := m
	m2["coffee"] = []string{"COffee"}
	m["tea"] = []string{"HOt tea"}
	fmt.Println(m)
	fmt.Println(m2)

}
