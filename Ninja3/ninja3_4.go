package main

import "fmt"

func Ninja4() {
	fmt.Println("Exercício 4:")
	for x := 2024; x >= 2003; x-- {
		idade := x - 2003
		fmt.Printf("Idade: %d\nAno: %d\n\n", idade, x)
	}
}
