package main

import "fmt"

func Ninja6() {
	fmt.Println("Exercício 6:")
	idade := 0
	for idade < 18 {
		fmt.Println("Menor de idade:", idade)
		idade++
	}
	if idade == 18 {
		fmt.Println("Adulto agora")
	}
}
