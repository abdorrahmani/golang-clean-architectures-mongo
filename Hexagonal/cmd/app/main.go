package main

import (
	"Hexagonal/configs"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := configs.LoadConfig()

	router := gin.Default()

	err := router.Run(cfg.ServerAddress)
	if err != nil {
		log.Print("Error starting server")
	}
}
