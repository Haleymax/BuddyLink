package main

import (
	"buddylink/config"
	"buddylink/internal/routers"
	"buddylink/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()

	config := config.LoadConfig()
	db, err := database.InitDB(config.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()
	routers.SetupRouter(r, config, db)

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
