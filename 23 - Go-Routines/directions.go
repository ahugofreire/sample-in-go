package main

import (
	"fmt"
)

func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}

// Channel receiver only
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// Channel send only
func ler(data <-chan string) {
	fmt.Println(<-data)
}
