package main

import (
	"Feature-Based/configs"
	"Feature-Based/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := configs.LoadConfig()

	mongoClient := db.ConnectMongo(config.MongoURI)
	defer mongoClient.Disconnect()

	r := gin.Default()

	err := r.Run(config.ServerAddress)
	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
