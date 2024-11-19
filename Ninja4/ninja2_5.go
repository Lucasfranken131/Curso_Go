package main

import "fmt"

func Ninja5() {
	fmt.Println("Exercício 5:")
	mapa := map[string][]string{
		"Ayrton Senna": []string{
			"Andar de carro", "Viajar", "Comprar carros",
		},
		"Celso Portiolli": []string{
			"Apresentar", "Pilotar avião", "11 de Setembro",
		},
		"Amelia Earhart": []string{
			"Pilotar", "Viajar", "Desbravar",
		},
	}

	for w, x := range mapa {
		fmt.Println(w)
		for y, z := range x {
			fmt.Printf("\t %d - %s\n", y, z)
		}
	}

	fmt.Println()
}
