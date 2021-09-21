package main

import (
	"awesomeProject/configs"
	"awesomeProject/server"
	"context"
	"github.com/DTB4/logger/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configs.NewConfig("configs/config.env")
	cfg := configs.InitConfig()
	myLogger := logger.NewLogger(cfg.LogsPath)
	srv := server.Start(cfg)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			myLogger.InfoLog("Server Message", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		myLogger.InfoLog("Shutdown info: ", err)
	}

}
