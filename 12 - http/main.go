package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://api.github.com/users/ahugofreire")
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close() // É importante fechar o body
}
