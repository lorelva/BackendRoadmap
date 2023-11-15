package cash

import "fmt"

type Cash struct {
	Amount float64
}

func (c Cash) Pay() {
	fmt.Printf(" %.2f", c.Amount)
}
