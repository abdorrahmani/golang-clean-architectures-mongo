package main

import (
	"ddd/configs"
	"ddd/internal/post"
	"ddd/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := configs.LoadConfig()

	mongoClient := db.ConnectMongo(config.MongoURI)
	defer mongoClient.Disconnect()
	db := mongoClient.GetDatabase(config.DBName)
	postRepo := post.NewPostRepository(db)
	postUC := post.NewPostUseCase(postRepo)
	postHandler := post.NewPostHandler(postUC)

	r := gin.Default()
	post.RegisterPostRoutes(r, postHandler)
	err := r.Run(config.ServerAddress)
	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
