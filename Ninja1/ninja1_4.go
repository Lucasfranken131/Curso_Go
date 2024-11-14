package main

import (
	"fmt"
)

func Ninja4() {
	type tipo int
	var x tipo
	fmt.Println("Exercício 4:")
	fmt.Printf("Tipo x: %T\n", x)
	fmt.Printf("Valor x: %d\n", x)
	x = 42
	fmt.Printf("Valor de x: %d\n", x)
	fmt.Println()
}
