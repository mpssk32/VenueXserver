package service_test

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"concert-service/internal/service"
	"testing"
)

func TestCreateVenue(t *testing.T) {

	config.ConnectDB()

	venue := models.Venue{
		Name:        "Arena",
		City:        "Amsterdam",
		Description: "Big venue",
		Equipment:   "Lights",
		Capacity:    1000,
		OwnerID:     "550e8400-e29b-41d4-a716-446655440000",
	}

	err := service.CreateVenue(venue)

	if err != nil {

		t.Log(err)
	}
}

func TestCreateVenue_InvalidData(t *testing.T) {

	config.ConnectDB()

	venue := models.Venue{}

	err := service.CreateVenue(venue)

	if err == nil {

		t.Error(
			"expected service error",
		)
	}
}

func TestCreateVenue_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	venue := models.Venue{
		Name:        "Arena",
		City:        "Amsterdam",
		Description: "Big venue",
		Equipment:   "Lights",
		Capacity:    1000,
		OwnerID:     "invalid",
	}

	err := service.CreateVenue(venue)

	if err == nil {

		t.Error(
			"expected uuid error",
		)
	}
}