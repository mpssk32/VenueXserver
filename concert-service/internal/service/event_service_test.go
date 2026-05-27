package service_test

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"concert-service/internal/service"
	"testing"
)

func TestCreateEvent(t *testing.T) {

	config.ConnectDB()

	event := models.Event{
		Title:    "Metallica Live",
		VenueID:  "550e8400-e29b-41d4-a716-446655440000",
		ArtistID: "550e8400-e29b-41d4-a716-446655440001",
	}

	err := service.CreateEvent(event)

	if err != nil {

		t.Log(err)
	}
}

func TestGetAllEvents(t *testing.T) {

	config.ConnectDB()

	_, err := service.GetAllEvents()

	if err != nil {

		t.Log(err)
	}
}

func TestSearchEvents(t *testing.T) {

	config.ConnectDB()

	_, err := service.SearchEvents(
		"Metallica",
	)

	if err != nil {

		t.Log(err)
	}
}

func TestGetEventByID_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	_, err := service.GetEventByID(
		"invalid",
	)

	if err == nil {

		t.Error(
			"expected uuid error",
		)
	}
}

func TestCreateEvent_InvalidData(t *testing.T) {

	config.ConnectDB()

	event := models.Event{}

	err := service.CreateEvent(event)

	if err == nil {

		t.Error(
			"expected service error",
		)
	}
}