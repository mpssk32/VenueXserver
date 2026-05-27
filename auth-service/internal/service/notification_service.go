package service

import (
	"auth-service/internal/config"
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishNotification(message string) error {

	if config.RabbitChannel == nil {

		return nil
	}

	return config.RabbitChannel.PublishWithContext(
		context.Background(),
		"",
		"notifications",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}