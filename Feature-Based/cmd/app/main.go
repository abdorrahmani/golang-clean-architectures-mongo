package main

import (
	"Feature-Based/configs"
	"Feature-Based/internal/post/handler"
	"Feature-Based/internal/post/repository"
	"Feature-Based/internal/post/service"
	"Feature-Based/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := configs.LoadConfig()

	mongoClient := db.ConnectMongo(config.MongoURI)
	defer mongoClient.Disconnect()

	database := mongoClient.GetDatabase(config.DBName)
	postRepo := repository.NewPostRepository(database)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := gin.Default()
	postHandler.RegisterRoutes(r)

	err := r.Run(config.ServerAddress)
	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
