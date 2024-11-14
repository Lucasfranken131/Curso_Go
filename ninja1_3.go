package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

func Ninja3() {
	fmt.Println("Exercício 3:")
	s := fmt.Sprint(a, " ", b, " ", c)
	fmt.Println("Valor de s:", s)
	fmt.Println()
}
