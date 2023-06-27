package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValor(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	conta := &Conta{
		Saldo:   100.0,
		Titular: "João",
	}

	adicionarValor(conta, 50.0)
	fmt.Println(conta.Saldo)
}
