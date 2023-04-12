package main

import "fmt"

func main() {
	canal := make(chan string) //canal vazio

	//Thread 2
	go func() {
		canal <- "ola mundo" // canal preenchido(cheio)
	}()

	//Thread 1
	msg := <-canal //canal esvazia
	fmt.Println(msg)
}
