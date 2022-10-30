package main

import "fmt"

// Escopo Global
const nome = "Hugo"

var (
	a bool
	b string
	c int
	d float64
)

var str string = "teste"

func main() {
	fmt.Println(nome)
	fmt.Println(a)   //false
	fmt.Println(b)   //""
	fmt.Println(c)   //0
	fmt.Println(d)   //0
	fmt.Println(str) //teste
}
