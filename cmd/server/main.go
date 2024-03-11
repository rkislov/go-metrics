package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rkislov/go-metrics.git/internal/config"
	"github.com/rkislov/go-metrics.git/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	dataServer := server.New(*cfg.Server)
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-cancelChan
		cancel()
	}()
	dataServer.Run(ctx)

	log.Println("Program end")
}
