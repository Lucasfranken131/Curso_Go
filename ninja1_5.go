package main

import (
	"fmt"
)

func Ninja5() {
	type tipo int
	var x tipo
	fmt.Println("Exercício 5:")
	x = 42
	var y = int(x)
	fmt.Printf("Tipo y: %T\n", y)
	fmt.Printf("Valor y: %d\n", y)
	fmt.Println()
}
