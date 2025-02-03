package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConsumerService struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	simpUserRepo *repositories.SimpUserRepository
}

type UserEvent struct {
	Type     string `json:"type"`
	SimpUser struct {
		ID           string `json:"id"`
		Email        string `json:"email"`
		IsChef       bool   `json:"isChef"`
		PartnerEmail string `json:"partnerEmail"`
	} `json:"user"`
}

func NewConsumerService(amqpURL string, simpUserRepo *repositories.SimpUserRepository) (*ConsumerService, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &ConsumerService{
		conn:         conn,
		channel:      channel,
		simpUserRepo: simpUserRepo,
	}, nil
}

func (s *ConsumerService) Start() error {
	// Declare exchange
	err := s.channel.ExchangeDeclare(
		"user_events", // name
		"fanout",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	// Declare queue
	queue, err := s.channel.QueueDeclare(
		"backend_user_events", // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		return err
	}

	// Bind queue to exchange
	err = s.channel.QueueBind(
		queue.Name,    // queue name
		"",            // routing key
		"user_events", // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := s.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	go s.handleMessages(msgs)
	return nil
}

func (s *ConsumerService) handleMessages(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		var event UserEvent
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		userID, _ := primitive.ObjectIDFromHex(event.SimpUser.ID)

		simpUser := &models.SimpUser{
			ID:           userID,
			Email:        event.SimpUser.Email,
			IsChef:       event.SimpUser.IsChef,
			PartnerEmail: event.SimpUser.PartnerEmail,
		}

		switch event.Type {
		case "user.created", "user.updated":
			if err := s.simpUserRepo.Upsert(context.Background(), simpUser); err != nil {
				log.Printf("Error upserting user: %v", err)
			}
		default:
			log.Printf("Unknown event type: %s", event.Type)
		}
	}
}

func (s *ConsumerService) Close() {
	if s.channel != nil {
		s.channel.Close()
	}
	if s.conn != nil {
		s.conn.Close()
	}
}
