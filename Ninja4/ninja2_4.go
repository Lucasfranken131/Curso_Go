package main

import "fmt"

func Ninja4() {
	fmt.Println("Exercício 4:")
	slico := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slicoP := &slico
	*slicoP = append(*slicoP, 52)
	*slicoP = append(*slicoP, 53, 54, 55, 56)
	// slico = append(*slicoP, 52)
	fmt.Println(*slicoP)
	fmt.Println()
}
