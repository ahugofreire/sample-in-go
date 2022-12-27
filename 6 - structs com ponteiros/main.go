package main

import "fmt"

type Cliente struct {
	nome string
}

// Passa uma copia
// func (c Cliente) andou() {
// Passar o valor na menoria
func (c *Cliente) andou() {
	c.nome = "John Doe"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	j := Cliente{
		nome: "John",
	}
	j.andou()
}
