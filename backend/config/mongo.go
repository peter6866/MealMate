package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI(AppConfig.MONGO_URI)

	// Set connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Create a MongoDB client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error while connecting to MongoDB: %v", err)
	}

	// Ping the db to check if the connection is successful
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("Error while pinging to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client
}