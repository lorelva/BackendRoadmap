package main

//FOLLOWING THE PRINCIPLE OCP :
//keep the existing code closed for modification and open
//for extension in interface PaymentGateway

type PaymentGateway interface {
	Pay(amount float64) bool
}

type CreditCardGateway struct {
}

func (g *CreditCardGateway) Pay(amount float64) bool {
	// Credit card payment logic
	return true
}

type WalletGateway struct {
}

func (g *WalletGateway) Pay(amount float64) bool {
	// Digital wallet payment logic
	return true
}

func main() {

}
