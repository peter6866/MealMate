package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderEventConsumer handles consuming order events from RabbitMQ
type OrderEventConsumer struct {
	userService *UserService
	channel     *amqp.Channel
}

// OrderEvent represents the structure of incoming order events
type OrderEvent struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}

// NewOrderEventConsumer creates a new OrderEventConsumer
func NewOrderEventConsumer(userService *UserService, channel *amqp.Channel) *OrderEventConsumer {
	return &OrderEventConsumer{
		userService: userService,
		channel:     channel,
	}
}

// Start begins consuming order events
func (c *OrderEventConsumer) Start() error {
	if err := c.setupExchangeAndQueue(); err != nil {
		return fmt.Errorf("failed to setup exchange and queue: %w", err)
	}

	return c.consume()
}

func (c *OrderEventConsumer) setupExchangeAndQueue() error {
	// Declare the exchange
	if err := c.channel.ExchangeDeclare(
		"order_events", // exchange name
		"fanout",       // exchange type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	); err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare a queue
	q, err := c.channel.QueueDeclare(
		"auth_order_events", // queue name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind the queue to the exchange
	if err := c.channel.QueueBind(
		q.Name,         // queue name
		"",             // routing key
		"order_events", // exchange
		false,          // no-wait
		nil,            // arguments
	); err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	return nil
}

func (c *OrderEventConsumer) consume() error {
	msgs, err := c.channel.Consume(
		"auth_order_events", // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	if err != nil {
		return fmt.Errorf("failed to start consuming: %w", err)
	}

	go c.handleMessages(msgs)
	log.Println("Started consuming order events")
	return nil
}

func (c *OrderEventConsumer) handleMessages(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		if err := c.processMessage(msg); err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}

func (c *OrderEventConsumer) processMessage(msg amqp.Delivery) error {
	var orderEvent OrderEvent
	if err := json.Unmarshal(msg.Body, &orderEvent); err != nil {
		return fmt.Errorf("failed to unmarshal order event: %w", err)
	}

	if orderEvent.Type != "order.created" {
		return nil // ignore other event types
	}

	objectID, err := primitive.ObjectIDFromHex(orderEvent.UserID)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %w", err)
	}

	if err := c.userService.repo.ClearCart(context.Background(), objectID); err != nil {
		return fmt.Errorf("failed to clear cart for user %s: %w", orderEvent.UserID, err)
	}

	return nil
}
