package main

import "fmt"

func main() {
	//For Default
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//For with range
	nums := []string{"um", "dois", "tres"}
	for key, value := range nums {
		fmt.Println(key, value)
	}

	// For type while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//loop
	for {
		fmt.Println("loop infinito!!")
	}
}
