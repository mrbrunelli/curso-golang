package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	fmt.Println("O valor de x é", x)

	// formatar o float para string para ser usado em concatenação
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	// %.2f força a retornar um float de 2 casas decimais apenas
	// %f retorna float
	// %s retorna string
	// %d retorna int
	// %t retorna boolean
	// %v detecta o tipo e retorna de acordo
	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v %v", a, b, b, c, d)
}
