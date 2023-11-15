package card

import "fmt"

type DebitCard struct {
	Amount float64
}

func (d DebitCard) Pay() {
	fmt.Printf(" %.2f", d.Amount)
}
