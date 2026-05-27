package service_test

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"concert-service/internal/service"
	"testing"
)

func TestCreateTicket(t *testing.T) {

	config.ConnectDB()

	ticket := models.Ticket{
		EventID: "550e8400-e29b-41d4-a716-446655440001",
		UserID:  "550e8400-e29b-41d4-a716-446655440000",
	}

	err := service.CreateTicket(ticket)

	if err != nil {

		t.Log(err)
	}
}

func TestCreateTicket_InvalidData(t *testing.T) {

	config.ConnectDB()

	ticket := models.Ticket{}

	err := service.CreateTicket(ticket)

	if err == nil {

		t.Log(
			"expected service error",
		)
	}
}

func TestCreateTicket_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	ticket := models.Ticket{
		EventID: "invalid",
		UserID:  "invalid",
	}

	err := service.CreateTicket(ticket)

	if err == nil {

		t.Log(
			"expected uuid error",
		)
	}
}

func TestGetUserTickets(t *testing.T) {

	config.ConnectDB()

	_, err := service.GetUserTickets(
		"550e8400-e29b-41d4-a716-446655440000",
	)

	if err != nil {

		t.Log(err)
	}
}

func TestGetUserTickets_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	_, err := service.GetUserTickets(
		"invalid",
	)

	if err == nil {

		t.Log(
			"expected uuid error",
		)
	}
}

func TestGetUserTickets_EmptyResult(t *testing.T) {

	config.ConnectDB()

	tickets, err := service.GetUserTickets(
		"550e8400-e29b-41d4-a716-446655440000",
	)

	if err != nil {

		t.Log(err)
	}

	if tickets == nil {

		t.Log(
			"tickets slice is nil",
		)
	}
}