package main

import (
	"auth-service/config"
	"auth-service/routes"
	"log"
)

func main() {
	config.LoadConfig()

	config.ConnectMongoDB()
	client := config.MongoClient

	// Initialize Gin Router
	router := routes.SetupRouter(client)

	// start server
	port := "8081"

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
