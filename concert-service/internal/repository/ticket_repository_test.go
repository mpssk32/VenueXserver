package repository_test

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"concert-service/internal/repository"
	"testing"
)

func TestCreateTicket(t *testing.T) {

	config.ConnectDB()

	ticket := models.Ticket{
		EventID: "550e8400-e29b-41d4-a716-446655440001",
		UserID:  "550e8400-e29b-41d4-a716-446655440000",
	}

	err := repository.CreateTicket(ticket)

	if err != nil {

		t.Log(err)
	}
}

func TestGetUserTickets(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetUserTickets(
		"550e8400-e29b-41d4-a716-446655440000",
	)

	if err != nil {

		t.Log(err)
	}
}
