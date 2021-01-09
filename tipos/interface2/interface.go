package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// estou usando ponteiro pq preciso acessar e modificar a ferrari na memória
// se eu não usar ponteiro, eu receberei uma cópia da instância
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// quando estou trabalhando com uma variável do tipo da interface,
	// e preciso escrever alguma informação,
	// é necessário usar o referenciador (&)
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
