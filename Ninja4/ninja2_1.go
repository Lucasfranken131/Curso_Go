package main

import "fmt"

func Ninja1() {
	fmt.Println("Exercício 1:")
	var slico = [5]int{0, 10, 20, 30, 45}
	for x, y := range slico {
		fmt.Printf("Posição de array: %d\nValor: %d\n\n", x, y)
	}
	fmt.Printf("Array do tipo: %T\n", slico)
	fmt.Println()
}
