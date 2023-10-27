package main

import "fmt"

//O --> OPEN-CLOSED PRINCIPLE (OCP)
// STATEMENT: a struct should be open for extension but closed for modification

// EXAMPLE -->payment system that will be able to process credit card payments

// PaymentMethod interface to add new payments methods without editing Payment behavior
type PaymentMethod interface {
	Pay()
}

// Payment struct is open for extension and closed for modification
type Payment struct{}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

type CreditCard struct {
	Amount float64
}

func (c CreditCard) Pay() {
	fmt.Printf("Paid %.2f using CreditCard", c.Amount)

}

// Add a new payment method and function
type PayPal struct {
	Amount float64
}

func (pp PayPal) Pay() {
	fmt.Printf("\nPaid %.2f using PayPal", pp.Amount)
}

// Add a second payment method
type ApplePay struct {
	Amount float64
}

func (a ApplePay) Pay() {
	fmt.Printf("\nPaid %.2f using ApplePay", a.Amount)
}

func main() {
	p := Payment{}
	c := CreditCard{10000.00}
	p.Process(c)

	//adding paypal method
	pp := PayPal{100.33}
	p.Process(pp)

	//adding a second : apple pay method
	a := ApplePay{9000.90}
	p.Process(a)
}
