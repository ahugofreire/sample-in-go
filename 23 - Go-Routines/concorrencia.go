package main

import (
	"fmt"
	"net/http"
	"sync"
)

var number uint64 = 0

func main() {

	//Resolvedo problema de reace condition com lock
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()

		//Outra opcao e usar variavel atomica sem lock
		//atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Vc é o numero, %d!", number)))
	})
	http.ListenAndServe(":3000", nil)
}

//Comando para verificar se está acontecendo race codition no golang
// go run -race main.go

//Comando apache ab
// ab -n 10000 -c 100 http://localhost:3000/
