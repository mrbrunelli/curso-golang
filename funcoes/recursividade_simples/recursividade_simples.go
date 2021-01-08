package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado1 := fatorial(5)
	fmt.Println(resultado1)

	resultado2 := fatorial(10)
	fmt.Println(resultado2)
}
