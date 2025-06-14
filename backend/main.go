package main

import (
	"buddylink/config"
	"buddylink/internal/routers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()

	config := config.LoadConfig()

	routers.SetupRouter(r, config)

	go func() {
		if err := r.Run(config.Server.Port); err != nil {
			log.Fatalf("Failed to start server:%v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
