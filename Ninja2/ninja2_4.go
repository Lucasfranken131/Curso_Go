package main

import "fmt"

func Ninja4() {
	x := 2
	fmt.Println("Exercício 4:")
	fmt.Printf("%d, %b, %#x", x, x, x)
	fmt.Println()
	y := x << 1
	fmt.Printf("%d, %b, %#x", y, y, y)
	fmt.Println()
}
