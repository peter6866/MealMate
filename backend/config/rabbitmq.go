package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQChannel *amqp.Channel

func ConnectRabbitMQ() {
	conn, err := amqp.Dial(AppConfig.RABBITMQ_URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	// Declare the exchange
	err = ch.ExchangeDeclare(
		"user_events", // exchange name
		"fanout",      // exchange type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	RabbitMQChannel = ch
}
