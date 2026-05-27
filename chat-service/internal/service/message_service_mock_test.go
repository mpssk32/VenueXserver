package service_test

import (
	"chat-service/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) SaveMessage(
	message models.Message,
) error {

	args := m.Called(message)

	return args.Error(0)
}

func (m *MockMessageRepository) GetMessages(
	senderID string,
	receiverID string,
) ([]models.Message, error) {

	args := m.Called(
		senderID,
		receiverID,
	)

	return args.Get(0).([]models.Message),
		args.Error(1)
}