package tests

import (
	"chat-service/internal/config"
	"chat-service/internal/service"
	"testing"
)

func TestGetMessages(t *testing.T) {

	config.ConnectDB()

	_, err := service.GetMessages(
		"user1",
		"user2",
	)

	if err != nil {

		t.Errorf(
			"expected successful message retrieval, got %v",
			err,
		)
	}
}