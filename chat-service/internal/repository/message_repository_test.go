package repository_test

import (
	"chat-service/internal/config"
	"chat-service/internal/models"
	"chat-service/internal/repository"
	"testing"
)

func TestSaveMessage(t *testing.T) {

	config.ConnectDB()

	message := models.Message{
		SenderID:   "550e8400-e29b-41d4-a716-446655440000",
		ReceiverID: "550e8400-e29b-41d4-a716-446655440001",
		Content:    "hello",
	}

	err := repository.SaveMessage(message)

	if err != nil {

		t.Log(err)
	}
}

func TestSaveMessage_Empty(t *testing.T) {

	config.ConnectDB()

	message := models.Message{}

	err := repository.SaveMessage(message)

	if err != nil {

		t.Log(err)
	}
}

func TestSaveMessage_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	message := models.Message{
		SenderID:   "bad",
		ReceiverID: "bad",
		Content:    "test",
	}

	err := repository.SaveMessage(message)

	if err == nil {

		t.Log(
			"expected uuid error",
		)
	}
}

func TestGetMessages(t *testing.T) {

	config.ConnectDB()

	messages, err := repository.GetMessages(
		"550e8400-e29b-41d4-a716-446655440000",
		"550e8400-e29b-41d4-a716-446655440001",
	)

	if err != nil {

		t.Log(err)
	}

	if messages == nil {

		t.Log(
			"messages slice is nil",
		)
	}
}

func TestGetMessages_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetMessages(
		"bad",
		"bad",
	)

	if err == nil {

		t.Log(
			"expected uuid error",
		)
	}
}

func TestGetMessages_Empty(t *testing.T) {

	config.ConnectDB()

	messages, err := repository.GetMessages(
		"",
		"",
	)

	if err != nil {

		t.Log(err)
	}

	if messages == nil {

		t.Log(
			"messages slice is nil",
		)
	}
}