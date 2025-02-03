package services

import (
	"auth-service/config"
	"auth-service/models"
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

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

func PublishUserEvent(ctx context.Context, eventType string, user *models.User) error {
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
		return err
	}

	return config.RabbitMQChannel.PublishWithContext(
		ctx,
		"user_events",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
