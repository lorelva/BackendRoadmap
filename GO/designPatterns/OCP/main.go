package main

import (
	"github.com/lorelva/BackendRoadmap/GO/designPatterns/OCP/payment"
	"github.com/lorelva/BackendRoadmap/GO/designPatterns/OCP/payment/method/card"
	"github.com/lorelva/BackendRoadmap/GO/designPatterns/OCP/payment/method/cash"
)

func main() {
	p := payment.Payment{}
	c := card.CreditCard{
		Amount: 2324.53,
	}
	d := card.DebitCard{}
	d.Amount = 2456.78

	ca := cash.Cash{
		Amount: 23.78,
	}

	p.Process(c)
	p.Process(d)
	p.Process(ca)

}
