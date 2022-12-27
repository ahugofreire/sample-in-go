package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Olá, mundo!!"

	showType(x)
	showType(y)

	//TYPE Assertion
	fmt.Println(y.(string))
	res, ok := y.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}

func showType(t interface{}) {
	fmt.Printf("o tipo da variavel é %T e o valor é %v\n", t, t)
}
