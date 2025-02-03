package services

import (
	"context"
	"log"

	"github.com/peter6866/foodie/config"
	"github.com/peter6866/foodie/repositories"
)

func StartConsumer() {
	client := config.MongoClient

	simpUserRepo := repositories.NewSimpUserRepository(client)

	consumerService, err := NewConsumerService(config.AppConfig.RABBITMQ_URL, simpUserRepo)
	if err != nil {
		log.Fatalf("Failed to start consumer service: %v", err)
	}

	go func() {
		<-context.Background().Done()
		consumerService.Close()
	}()

	go func() {
		if err := consumerService.Start(); err != nil {
			log.Fatalf("Failed to start consumer service: %v", err)
		}
	}()

	log.Println("Consumer service started")
}
