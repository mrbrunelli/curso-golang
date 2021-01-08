package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func soma(a, b int) (total int) {
	total = a + b
	return
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado1 := exec(multiplicacao, 3, 4)
	fmt.Println(resultado1)

	resultado2 := exec(soma, 5, 10)
	fmt.Println(resultado2)
}
