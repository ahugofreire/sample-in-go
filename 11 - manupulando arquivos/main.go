package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file_test.txt") //vai criar arquivo.
	if err != nil {
		panic(err)
	}

	//tamanho, err := f.Write([]byte("Escrevendo dados no arquivos")) //escrever dados
	tamanho, err := f.WriteString("Hello, World!") //Escrever no arquivo
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	defer f.Close()

	//Leitura de arquivo
	arquivo, err := os.ReadFile("file_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//Leitura de pouco em pouco, carregando partes do arquivos
	arq, err := os.Open("file_test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arq)
	buffer := make([]byte, 5) //tamanho dos pedaços

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break //vai entra quando acabar o buffer
		}
		fmt.Println(string(buffer[:n]))
	}

	//Remover um arquivo
	err = os.Remove("file_test.txt")
	if err != nil {
		panic(err)
	}

}
