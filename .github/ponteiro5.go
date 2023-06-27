package main

import (
	"fmt"
	"math"
)

func médiaAritmetica(numero *float64) {
	*numero = (*numero + math.Pi) / 2
}

func main() {
	valor := 3.14
	médiaAritmetica(&valor)
	fmt.Println(valor) // Saída: 3.442
}
