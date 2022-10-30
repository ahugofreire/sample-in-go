package main

import "fmt"

func main() {
	fmt.Println(mult(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

// Exemplo de uso da func
func sum(a, b int) (int, bool) {
	if a+b <= 0 {
		return a + b, false
	}

	return a + b, true
}

// Exemplo de uso func variadicas, é possivel passar diversos valores daquele tipo
func mult(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

// Exemplo de uso func closore
func avg() {
	media := func() int {
		return mult(1, 2, 3, 4) / 2
	}()

	fmt.Println(media)
}
