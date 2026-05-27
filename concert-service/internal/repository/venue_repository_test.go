package repository_test

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"concert-service/internal/repository"
	"testing"
)

func TestCreateVenue(t *testing.T) {

	config.ConnectDB()

	venue := models.Venue{
		Name:        "Arena",
		City:        "Amsterdam",
		Description: "Big concert hall",
		Equipment:   "Lights",
		Capacity:    1000,
		OwnerID:     "550e8400-e29b-41d4-a716-446655440000",
	}

	err := repository.CreateVenue(venue)

	if err != nil {

		t.Log(err)
	}
}

func TestGetVenues(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetVenues()

	if err != nil {

		t.Log(err)
	}
}



func TestCreateVenue_InvalidData(t *testing.T) {

	config.ConnectDB()

	venue := models.Venue{}

	err := repository.CreateVenue(venue)

	if err == nil {

		t.Error(
			"expected repository error",
		)
	}
}