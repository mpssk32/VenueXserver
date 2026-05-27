package repository_test

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"concert-service/internal/repository"
	"testing"
)

func TestCreateApplication(t *testing.T) {

	config.ConnectDB()

	app := models.Application{
		ArtistID: "550e8400-e29b-41d4-a716-446655440000",
		VenueID:  "550e8400-e29b-41d4-a716-446655440001",
		Message:  "Hello venue",
		Status:   "pending",
	}

	err := repository.CreateApplication(app)

	if err != nil {

	t.Log(err)
}
}

func TestGetVenueApplications(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetVenueApplications(
		"550e8400-e29b-41d4-a716-446655440001",
	)

		if err != nil {

		t.Log(err)
	}
}

func TestUpdateApplicationStatus(t *testing.T) {

	config.ConnectDB()

	err := repository.UpdateApplicationStatus(
		"1",
		"550e8400-e29b-41d4-a716-446655440001",
		"approved",
	)

	if err != nil {

		t.Log(err)
	}
}

func TestDeleteApplication(t *testing.T) {

	config.ConnectDB()

	err := repository.DeleteApplication(
		"1",
	)

	if err != nil {

	t.Log(err)
}
}
