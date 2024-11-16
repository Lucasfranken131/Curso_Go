package main

import "fmt"

func Ninja6() {
	fmt.Println("Exercício 6:")
	const (
		a = iota + 2024
		b
		c
		d
		e
	)
	fmt.Println(a, b, c, d, e)
	fmt.Println()
}
