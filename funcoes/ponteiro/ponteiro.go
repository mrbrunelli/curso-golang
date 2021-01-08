package main

import "fmt"

// por padrão o Go trabalha com cópias
// ou seja, quando passamos uma variavel por parâmetro, ele está passando uma cópia, e não uma referência

func inc1(n int) {
	n++
}

// resvisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// resvisão: & é usado para obetr o endereço da variável
	inc2(&n) // por referência
	fmt.Println(n)
}
