package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func CreateConcert(concert models.Concert) error {
	return repository.CreateConcert(concert)
}

func GetConcerts() ([]models.Concert, error) {
	return repository.GetConcerts()
}