package service

import (
	"chat-service/internal/models"
	"chat-service/internal/repository"
)

func SaveMessage(message models.Message) error {
	return repository.SaveMessage(message)
}

func GetMessages(
	senderID string,
	receiverID string,
) ([]models.Message, error) {

	return repository.GetMessages(
		senderID,
		receiverID,
	)
}