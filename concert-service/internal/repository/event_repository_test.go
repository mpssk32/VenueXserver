package repository_test

import (
	"concert-service/internal/config"
	"concert-service/internal/repository"
	"testing"
)

func TestSearchEventsRepository(t *testing.T) {

	config.ConnectDB()

	_, err := repository.SearchEvents(
		"Metallica",
	)

	if err != nil {

		t.Errorf(
			"expected successful repository search, got %v",
			err,
		)
	}
}

func TestGetEventByID_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetEventByID(
		"invalid",
	)

	if err == nil {

		t.Error(
			"expected error for invalid uuid",
		)
	}
}


func TestSearchEvents_Empty(t *testing.T) {

	config.ConnectDB()

	_, err := repository.SearchEvents("")

	if err != nil {

		t.Log(err)
	}
}