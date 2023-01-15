package main

import "github.com/ahugofreire/22/apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}
