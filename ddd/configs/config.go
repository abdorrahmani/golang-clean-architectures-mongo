package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	MongoURI      string
	DBName        string
	ServerAddress string
}

func LoadConfig() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		MongoURI:      os.Getenv("MONGO_URI"),
		DBName:        os.Getenv("DB_NAME"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}
}
