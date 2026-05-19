package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func CreateVenue(venue models.Venue) error {
	return repository.CreateVenue(venue)
}