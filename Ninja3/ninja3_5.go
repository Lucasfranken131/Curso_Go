package main

import "fmt"

func Ninja5() {
	fmt.Println("Exercício 5:")
	for z := 10; z <= 40; z++ {
		fmt.Printf("%d dividido por 4 resta: %d\n", z, (z % 4))
	}
	fmt.Println()
}
