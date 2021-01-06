package main

// para usar alias no import, é só passar dessa forma: import m "math"
// para ignorar um import, é só colocar dessa forma: import _ "math"
import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 inferido pelo compilador

	area := PI * math.Pow(raio, 2) // forma reduzida de criar uma var
	fmt.Println("A área da circuferência é", area)

	// outra forma de declarar constantes e variáveis
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// atribuir variáveis em linha
	var e, f bool = true, false
	fmt.Println(e, f)

	// forma reduzida em linha
	g, h, i := 2, false, "opa!"
	fmt.Println(g, h, i)
}
