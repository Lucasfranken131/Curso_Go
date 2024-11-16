package main

import "fmt"

func Ninja9() {
	fmt.Println("Exercício 9:")
	esporteFavorito := "Handebol"
	switch {
	case esporteFavorito == "Futebol":
		fmt.Println("Você chuta bola!")
	case esporteFavorito == "Vôlei":
		fmt.Println("Você empurra bola!")
	case esporteFavorito == "Basquete":
		fmt.Println("Você joga bola!")
	default:
		fmt.Println("Você quer ser especial!")
	}
	fmt.Println()
}
