package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int     `json:"-"` //ignorar propiedade do json
	Saldo  float64 `json:"saldo"`
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  100.0,
	}
	result, err := json.Marshal(conta) //converter para Json
	if err != nil {
		println(err)
	}
	println(string(result))

	//Converter e já retorna
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	//Converter de json para struct
	jsonPuro := []byte(`{"Numero": 2,"Saldo": 220.0}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Numero)
}
