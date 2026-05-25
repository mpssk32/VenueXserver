package tests

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
		VenueID:  "venue123",
		ArtistID: "artist123",
	}

	err := service.CreateEvent(event)

	if err != nil {

		t.Errorf(
			"expected successful event creation, got %v",
			err,
		)
	}
}


func TestSearchEvents(t *testing.T) {

	config.ConnectDB()

	_, err := service.SearchEvents(
		"Metallica",
	)

	if err != nil {

		t.Errorf(
			"expected successful search, got %v",
			err,
		)
	}
}


func TestGetEventByID(t *testing.T) {

	config.ConnectDB()

	_, err := service.GetEventByID(
		"1",
	)

	if err != nil {

		t.Errorf(
			"expected event lookup, got %v",
			err,
		)
	}
}