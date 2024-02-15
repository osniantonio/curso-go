package main

import "github.com/osniantonio/goexpert/15-APIS/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
