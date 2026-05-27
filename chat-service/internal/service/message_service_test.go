package service_test

import (
	"chat-service/internal/models"
	"chat-service/internal/service"
	"errors"
	"testing"
)

func TestSaveMessage(t *testing.T) {

	mockRepo := new(MockMessageRepository)

	svc := service.NewMessageService(
		mockRepo,
	)

	message := models.Message{
		SenderID:   "1",
		ReceiverID: "2",
		Content:    "hello",
	}

	mockRepo.
		On("SaveMessage", message).
		Return(nil)

	err := svc.SaveMessage(message)

	if err != nil {

		t.Errorf(
			"unexpected error %v",
			err,
		)
	}

	mockRepo.AssertExpectations(t)
}

func TestSaveMessage_Error(t *testing.T) {

	mockRepo := new(MockMessageRepository)

	svc := service.NewMessageService(
		mockRepo,
	)

	message := models.Message{}

	mockRepo.
		On("SaveMessage", message).
		Return(errors.New("db error"))

	err := svc.SaveMessage(message)

	if err == nil {

		t.Errorf(
			"expected error",
		)
	}

	mockRepo.AssertExpectations(t)
}

func TestGetMessages(t *testing.T) {

	mockRepo := new(MockMessageRepository)

	svc := service.NewMessageService(
		mockRepo,
	)

	expected := []models.Message{
		{
			SenderID: "1",
			Content:  "hello",
		},
	}

	mockRepo.
		On(
			"GetMessages",
			"1",
			"2",
		).
		Return(expected, nil)

	messages, err := svc.GetMessages(
		"1",
		"2",
	)

	if err != nil {

		t.Errorf(
			"unexpected error %v",
			err,
		)
	}

	if len(messages) != 1 {

		t.Errorf(
			"expected 1 message",
		)
	}

	mockRepo.AssertExpectations(t)
}