package main

import "fmt"

func Ninja1() {
	fmt.Println("Exercício 1:")
	type pessoa struct {
		nome              string
		sobrenome         string
		sabores_favoritos []string
	}

	pessoa1 := pessoa{
		nome:              "Pelé",
		sobrenome:         "Silva",
		sabores_favoritos: []string{"napolitano", "chocolate"},
	}

	pessoa2 := pessoa{
		nome:              "Jussara",
		sobrenome:         "Costa",
		sabores_favoritos: []string{"Baunilha", "morango", "chocolate", "açaí"},
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	fmt.Println()
}
