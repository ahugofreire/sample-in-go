package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	// o Select vai imprimir o canal que chagar a informação primerio;
	// Caso passe 4 segundo e não chagar a inforção em nenhum canal da timeout
	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)

	case msg2 := <-c2:
		fmt.Println("received", msg2)

	case <-time.After(time.Second * 4):
		fmt.Println("Timeout")
	}

}
