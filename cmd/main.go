package main

import (
	"awesomeProject/configs"
	"awesomeProject/server"
)

func main() {
	configs.NewConfig("configs/config.env")
	cfg := configs.InitConfig()
	server.Start(cfg)

}
