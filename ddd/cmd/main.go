package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
