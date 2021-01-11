package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println("Executou!") // pode ser que ele consiga ler essa linha ainda, mas se eu mandar pra baixo, ele não conseguirá mais
	c <- 5
	c <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
