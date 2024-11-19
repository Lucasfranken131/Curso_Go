package main

import "fmt"

func Ninja2() {
	fmt.Println("Exercício 2:")
	var slico = []int{0, 10, 20, 30, 45, 50, 55, 29, 300, 1}
	for x, y := range slico {
		fmt.Printf("Posição de array: %d\nValor: %d\n\n", x, y)
	}
	fmt.Printf("Array do tipo: %T\n", slico)
	fmt.Println()
}
