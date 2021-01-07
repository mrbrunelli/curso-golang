package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15457.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11345.45,
		},
		"P": {
			"Pedro Júnior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Printf("\nLetra: %s\n", letra)
		for nome, salario := range funcs {
			fmt.Printf("Nome: %s, Salário: %f\n", nome, salario)
		}
	}
}
