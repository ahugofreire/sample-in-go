package main

import "fmt"

func main() {
	salary := map[string]int64{"john": 1000, "bob": 1500, "alice": 3000, "lee": 4500}
	fmt.Println(salary["john"]) //mostrando chave
	delete(salary, "bob")       //Removendo chave
	salary["alice"] = 3800      //trocando valor

	//Outra forma de iniciliar o map
	sal := make(map[string]int)
	sal["wes"] = 8000

	//Percorer um map
	for name, s := range salary {
		fmt.Printf("O salario de %s é %d\n", name, s)
	}
}
