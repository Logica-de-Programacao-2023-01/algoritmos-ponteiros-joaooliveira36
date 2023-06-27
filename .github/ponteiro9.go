package main

import (
	"fmt"
)

type Livros struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livros) {
	desconto := livro.Preco * 0.1
	livro.Preco -= desconto
}

func main() {
	livro := &Livros{
		Titulo: "Aventuras de Alice no País das Maravilhas",
		Autor:  "Lewis Carroll",
		Preco:  29.99,
	}

	fmt.Println("Preço antes do desconto:", livro.Preco)

	aplicarDesconto(livro)

	fmt.Println("Preço após o desconto:", livro.Preco)
}
