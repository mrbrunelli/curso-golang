package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[79879843292] = "Pedro"
	aprovados[99432429872] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[99432429872])
	delete(aprovados, 99432429872)
	fmt.Println(aprovados[99432429872])
}
