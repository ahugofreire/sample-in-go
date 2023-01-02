package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	//TimeOut
	c := http.Client{Timeout: time.Second}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	//Buffer

	b := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "wesley"}`))
	response, err := b.Post("http://ge.com.br", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	io.CopyBuffer(os.Stdout, response.Body, nil)
}
