package service

import (
	"concert-service/internal/models"
	"concert-service/internal/repository"
)

func CreateConcert(concert models.Concert) error {
	return repository.CreateConcert(concert)
}

func GetConcerts() ([]models.Concert, error) {
	return repository.GetConcerts()
}