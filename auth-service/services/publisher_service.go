package services

import (
	"auth-service/models"
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// UserEventPublisher handles publishing user events to RabbitMQ
type UserEventPublisher struct {
	channel *amqp.Channel
}

type UserEvent struct {
	Type string   `json:"type"`
	User SimpUser `json:"user"`
}

type SimpUser struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	IsChef       bool   `json:"isChef"`
	PartnerEmail string `json:"partnerEmail"`
}

// NewUserEventPublisher creates a new UserEventPublisher
func NewUserEventPublisher(channel *amqp.Channel) *UserEventPublisher {
	return &UserEventPublisher{
		channel: channel,
	}
}

// PublishUserEvent publishes a user event to RabbitMQ
func (p *UserEventPublisher) PublishUserEvent(ctx context.Context, eventType string, user *models.User) error {
	event := UserEvent{
		Type: eventType,
		User: SimpUser{
			ID:           user.ID.Hex(),
			Email:        user.Email,
			IsChef:       user.IsChef,
			PartnerEmail: user.PartnerEmail,
		},
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal user event: %w", err)
	}

	if err := p.channel.PublishWithContext(
		ctx,
		"user_events", // exchange
		"",            // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		return fmt.Errorf("failed to publish user event: %w", err)
	}

	return nil
}
