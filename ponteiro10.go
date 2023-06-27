package main

import (
	"fmt"
)

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}
	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}
func preencherPrimos(primos *[]int, tamanho int) {
	count := 0
	numero := 2
	for count < tamanho {
		if isPrimo(numero) {
			*primos = append(*primos, numero)
			count++
		}
		numero++
	}
}
func main() {
	var primos []int
	tamanho := 10

	preencherPrimos(&primos, tamanho)

	fmt.Println(primos)
}
