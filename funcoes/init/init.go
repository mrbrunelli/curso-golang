package main

import "fmt"

func init() { // executa antes do método main
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
