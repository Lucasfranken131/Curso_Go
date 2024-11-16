package main

import "fmt"

func Ninja3() {
	fmt.Println("Exercício 3:")
	ano := 2024
	for ano >= 2003 {
		idade := ano - 2003
		fmt.Printf("Idade: %d\nAno: %d\n\n", idade, ano)
		ano -= 1
	}
}
