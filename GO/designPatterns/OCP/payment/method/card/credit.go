package card

import "fmt"

type CreditCard struct {
	Amount float64
}

func (c CreditCard) Pay() {
	fmt.Printf(" %.2f", c.Amount)
	//abierto a extensión agregando más código sin afectar a otros
	for i := 0; i < 5; i++ {
		fmt.Println("Hola")
	}
}
