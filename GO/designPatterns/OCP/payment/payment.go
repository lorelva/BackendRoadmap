package payment

import (
	"fmt"
	"time"
)

type PaymentMethod interface {
	Pay()
}

type Payment struct {
	Date time.Time
	Name string
}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
	fmt.Println("Pago finalizado")
}
