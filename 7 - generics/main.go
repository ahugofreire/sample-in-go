package main

import "fmt"

type Mynumber int

// o ~ seria para usar o retorno do MyNumber
type Number interface {
	~int | float64
}

// ambas as formas de usar o generics
// func Soma[T int | float64](m map[string]T) T {
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Comparação
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	f := map[string]int{"Wesley": 10.00, "João": 20.00, "Maria": 300.0}
	fmt.Println(Soma(m))
	fmt.Println(Soma(f))

	m2 := map[string]Mynumber{"Wesley": 10.00, "João": 20.00, "Maria": 300.0}
	fmt.Println(Soma(m2))

	fmt.Println(Compara(10, 10))

	//mais info pacote https://pkg.go.dev/golang.org/x/exp/constraints
}
