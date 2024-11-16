package main

import "fmt"

func Ninja2() {
	fmt.Println("Exercício 2:")
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for y := 0; y < 3; y++ {
			fmt.Printf("%#U\n", x)
		}
	}
	fmt.Println()
}
