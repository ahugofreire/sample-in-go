package main

func main() {
	//Memoria -> endereço -> valor
	a := 10
	println(a)  //10
	println(&a) //0xc00004e768

	// o * sempre aponta para o endereço que está na memoria
	var ponteiro *int = &a
	println(ponteiro) //0xc00004e768
	*ponteiro = 20
	println(a) //20

	b := &a
	println(b)  //0xc00004e768
	println(*b) //20

	var1 := 10
	var2 := 20
	println(sum(&var1, &var2))

	//O ponteiro deve ser usado para trabalhar com valores mutaveis
}

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}
