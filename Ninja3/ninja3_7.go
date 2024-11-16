package main

import "fmt"

func Ninja7() {
	fmt.Println("Exercício 7:")
	idade := 0
	for idade <= 20 {
		if idade >= 18 {
			fmt.Println("Adulto agora:", idade)
		} else {
			fmt.Println("Menor de idade:", idade)
		}
		idade++
	}
	fmt.Println()
}
