package main

import "fmt"

func num_par(numero *int) {
	if *numero%2 == 0 {
		*numero = 0
	} else {
		*numero = 1
	}
}

func main() {
	numero := 5
	num_par(&numero)
	fmt.Print(numero)
}
