package main

import "fmt"

func atualizarValor(p *int) {
	numero := *p

	ultimoDigito := numero % 10
	numero = numero / 10
	penultimoDigito := numero % 10

	*p = ultimoDigito + penultimoDigito
}

func main() {
	valor := 1234
	atualizarValor(&valor)
	fmt.Println(valor)
}
