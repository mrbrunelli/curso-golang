package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//array1 = append(slice1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6) // append o slice cresce conforme for adicionado mais itens
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // copy não aumenta o tamanho do slice, ele copia só o que cabe dentro do slice
	fmt.Println(slice2)
}
