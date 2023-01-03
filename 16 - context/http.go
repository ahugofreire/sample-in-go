package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Panicln("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line stdout
		log.Panicln("Request processada com sucesso")
		//imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)
	}
}
