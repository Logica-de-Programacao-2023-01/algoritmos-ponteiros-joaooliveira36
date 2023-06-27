package main

import (
	"fmt"
)

func reverterString(p *string) {
	bytes := []byte(*p)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	*p = string(bytes)
}

func main() {
	texto := "Olá"
	reverterString(&texto)
	fmt.Println(texto)
}
