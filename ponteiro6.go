package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func atualizarLivro(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}
func main() {
	livro := Livro{
		Titulo: "Livro Exemplo",
		Autor:  "Anônimo",
	}
	atualizarLivro(&livro)
	fmt.Println(livro.Titulo)
}
