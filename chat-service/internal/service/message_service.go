package service

import (
	"chat-service/internal/models"
	"chat-service/internal/repository"
)

type MessageRepository interface {
	SaveMessage(message models.Message) error
	GetMessages(
		senderID string,
		receiverID string,
	) ([]models.Message, error)
}

type messageService struct {
	repo MessageRepository
}

func NewMessageService(
	repo MessageRepository,
) *messageService {

	return &messageService{
		repo: repo,
	}
}

func (s *messageService) SaveMessage(
	message models.Message,
) error {

	return s.repo.SaveMessage(message)
}

func (s *messageService) GetMessages(
	senderID string,
	receiverID string,
) ([]models.Message, error) {

	return s.repo.GetMessages(
		senderID,
		receiverID,
	)
}

type repositoryWrapper struct{}

func (repositoryWrapper) SaveMessage(
	message models.Message,
) error {

	return repository.SaveMessage(message)
}

func (repositoryWrapper) GetMessages(
	senderID string,
	receiverID string,
) ([]models.Message, error) {

	return repository.GetMessages(
		senderID,
		receiverID,
	)
}

var MessageService = NewMessageService(
	repositoryWrapper{},
)