package db

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

type MongoClient struct {
	Client *mongo.Client
}

func ConnectMongo(uri string) *MongoClient {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB successfully")
	return &MongoClient{Client: client}
}

func (mc *MongoClient) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := mc.Client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect MongoDB: %v", err)
	}
	log.Println("Disconnected MongoDB successfully")
}

func (mc *MongoClient) GetDatabase(name string) *mongo.Database {
	return mc.Client.Database(name)
}
