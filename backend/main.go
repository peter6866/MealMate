package main

import (
	"log"

	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/routes"
)

func main() {
	config.LoadConfig()

	config.ConnectMongoDB()
	client := config.MongoClient

	// Initialize Gin Router
	router := routes.SetupRouter(client)

	// start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
