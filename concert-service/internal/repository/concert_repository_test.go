package repository_test

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"concert-service/internal/repository"
	"testing"
)

func TestCreateConcert(t *testing.T) {

	config.ConnectDB()

	concert := models.Concert{
		Title:       "Metallica Live",
		Description: "Big concert",
		Genre:       "Metal",
		VenueID:     "550e8400-e29b-41d4-a716-446655440000",
		CreatedBy:   "550e8400-e29b-41d4-a716-446655440000",
		TicketPrice: 100,
	}

	err := repository.CreateConcert(concert)

	if err != nil {

		t.Log(err)
	}
}

func TestGetConcerts(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetConcerts()

	if err != nil {

		t.Log(err)
	}
}

func TestCreateConcert_InvalidData(t *testing.T) {

	config.ConnectDB()

	concert := models.Concert{}

	err := repository.CreateConcert(concert)

	if err == nil {

		t.Error(
			"expected repository error",
		)
	}
}

func TestGetConcerts_EmptyResult(t *testing.T) {

	config.ConnectDB()

	concerts, err := repository.GetConcerts()

	if err != nil {

		t.Log(err)
	}

	if concerts == nil {

		t.Log(
			"empty concerts list",
		)
	}
}

func TestCreateConcert_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	concert := models.Concert{
		Title:       "Concert",
		Description: "Desc",
		Genre:       "Rock",
		VenueID:     "invalid",
		CreatedBy:   "invalid",
		TicketPrice: 100,
	}

	err := repository.CreateConcert(concert)

	if err == nil {

		t.Error(
			"expected uuid error",
		)
	}
}

