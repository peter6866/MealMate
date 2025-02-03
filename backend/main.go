package main

import (
	"log"
	"os"

	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/routes"
	"github.com/peter6866/foodie/services"
)

func main() {
	config.LoadConfig()

	config.ConnectMongoDB()

	services.StartConsumer()

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
