package main

import (
	"awesomeProject/configs"
	"awesomeProject/server"
)

func main() {
	configs.NewConfig(false)
	cfg := configs.InitConfig()
	server.Start(cfg)

}
