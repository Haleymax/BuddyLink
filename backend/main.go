package main

import (
	"buddylink/config"
	"buddylink/internal/routers"
	"buddylink/pkg/cache_client"
	"buddylink/pkg/database"
	"buddylink/pkg/object_storage"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()
	config := config.GetConfig()
	db, err := database.InitDB(config.MySQL)
	defer database.CloseDB()

	err = cache_client.NewRedisClient(config.Redis)
	if err != nil {
		log.Println("redis init err:", err)
		panic(err)
	}

	defer func() {
		err := cache_client.Close()
		if err != nil {
			log.Println("redis close err:", err)
		}
	}()

	err = object_storage.NewMinio(config.Minio)
	if err != nil {
		log.Println("minio init err:", err)
		panic(err)
	}

	routers.SetupRouter(r, db)

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
