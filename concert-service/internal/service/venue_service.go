package service

import (
	"concert-service/internal/models"
	"concert-service/internal/repository"
)

func CreateVenue(venue models.Venue) error {
	return repository.CreateVenue(venue)
}