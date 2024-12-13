package main

import (
	"Hexagonal/configs"
	"Hexagonal/internal/adapters/api"
	"Hexagonal/internal/adapters/database"
	"Hexagonal/internal/core/post"
	"Hexagonal/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := configs.LoadConfig()

	mongoClient := db.ConnectMongo(cfg.MongoURI)
	defer mongoClient.Disconnect()

	getDatabase := mongoClient.GetDatabase(cfg.DBName)

	postRepo := database.NewPostRepository(getDatabase)
	postUseCase := post.NewPostUseCase(postRepo)
	postHandler := api.NewPostHandler(postUseCase)

	router := gin.Default()
	postHandler.RegisterRoutes(router)

	err := router.Run(cfg.ServerAddress)
	if err != nil {
		log.Print("Error starting server")
	}
}
