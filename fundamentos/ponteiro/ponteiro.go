package main

import "fmt"

func main() {
	i := 1 // não posso inicializar com nil uma variável que não é um ponteiro

	// Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor)
	i++

	// operações de aritimética não são permitidas em ponteiros ex: p++

	fmt.Println(p, *p, i, &i)
}
