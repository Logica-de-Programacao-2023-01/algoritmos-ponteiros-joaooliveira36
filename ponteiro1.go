package main

import (
	"fmt"
)

func atualizarSoma(p *int, n int) {
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	*p = soma
}

func main() {
	numero := 0
	n := 5
	atualizarSoma(&numero, n)
	fmt.Println(numero)
}
