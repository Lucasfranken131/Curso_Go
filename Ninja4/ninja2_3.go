package main

import "fmt"

func Ninja3() {
	fmt.Println("Exercício 3:")
	var slico = []int{0, 10, 20, 30, 45, 50, 55, 29, 300, 1}
	fmt.Println(slico[:3])
	fmt.Println(slico[5:])
	fmt.Println(slico[2:7])
	fmt.Println(slico[3 : len(slico)-1])
	fmt.Println()
}
