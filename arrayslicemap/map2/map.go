package main

import "fmt"

func main() {
	// chave e valor podem ser de qualquer tipo, desde que sejam homogêneos
	funcsESalarios := map[string]float64{
		"José João":        11325.45,
		"Gabriela Silva":   15456.78,
		"Pedro Júnior":     1200.0,
		"Matheus Brunelli": 4500.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
