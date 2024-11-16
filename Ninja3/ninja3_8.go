package main

import "fmt"

func Ninja8() {
	fmt.Println("Exercício 8:")
	integerme := 1010
	switch {
	case integerme == 0:
		fmt.Println("Integerme 0")
	case integerme == 1:
		fmt.Println("Integerme 1")
	case integerme == 2:
		fmt.Println("Integerme 2")
	default:
		fmt.Println("Integerme fora do número")
	}
	fmt.Println()
}
