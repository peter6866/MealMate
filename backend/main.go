package main

import (
	"log"
	"os"

	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/routes"
)

func main() {
	config.LoadConfig()
	config.ConnectMongoDB()
	config.ConnectRabbitMQ()

	// Initialize Gin Router
	router := routes.SetupRouter(config.MongoClient)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
