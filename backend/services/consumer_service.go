package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEventConsumer struct {
	simpUserRepo *repositories.SimpUserRepository
	channel      *amqp.Channel
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

func NewUserEventConsumer(simpUserRepo *repositories.SimpUserRepository, channel *amqp.Channel) *UserEventConsumer {
	return &UserEventConsumer{
		channel:      channel,
		simpUserRepo: simpUserRepo,
	}
}

func (c *UserEventConsumer) Start() error {
	if err := c.setupExchangeAndQueue(); err != nil {
		return fmt.Errorf("failed to setup exchange and queue: %w", err)
	}

	return c.consume()
}

func (c *UserEventConsumer) setupExchangeAndQueue() error {
	// Declare the exchange
	if err := c.channel.ExchangeDeclare(
		"user_events", // exchange name
		"fanout",      // exchange type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	); err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare a queue
	q, err := c.channel.QueueDeclare(
		"main_user_events", // queue name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind the queue to the exchange
	if err := c.channel.QueueBind(
		q.Name,        // queue name
		"",            // routing key
		"user_events", // exchange
		false,         // no-wait
		nil,           // arguments
	); err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	return nil
}

func (c *UserEventConsumer) consume() error {
	msgs, err := c.channel.Consume(
		"main_user_events", // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	go c.handleMessages(msgs)
	log.Println("Started consuming user events")
	return nil
}

func (c *UserEventConsumer) handleMessages(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		if err := c.processMessage(msg); err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}

func (c *UserEventConsumer) processMessage(msg amqp.Delivery) error {
	var userEvent UserEvent
	if err := json.Unmarshal(msg.Body, &userEvent); err != nil {
		return fmt.Errorf("failed to unmarshal user event: %w", err)
	}

	return c.handleUserEvent(context.Background(), &userEvent)
}

func (c *UserEventConsumer) handleUserEvent(ctx context.Context, event *UserEvent) error {
	userID, err := primitive.ObjectIDFromHex(event.SimpUser.ID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	switch event.Type {
	case "user.created", "user.updated":
		simpUser := &models.SimpUser{
			ID:           userID,
			Email:        event.SimpUser.Email,
			IsChef:       event.SimpUser.IsChef,
			PartnerEmail: event.SimpUser.PartnerEmail,
		}
		if err := c.simpUserRepo.Upsert(ctx, simpUser); err != nil {
			return fmt.Errorf("failed to upsert user: %w", err)
		}

	default:
		return fmt.Errorf("unknown event type: %s", event.Type)
	}

	return nil
}

func (c *UserEventConsumer) Close() error {
	if err := c.channel.Close(); err != nil {
		return fmt.Errorf("error closing channel: %w", err)
	}
	return nil
}
