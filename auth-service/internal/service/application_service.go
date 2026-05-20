package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func CreateApplication(app models.Application) error {
	return repository.CreateApplication(app)
}

func GetVenueApplications(ownerID string) ([]models.Application, error) {
	return repository.GetVenueApplications(ownerID)
}

func UpdateApplicationStatus(
	applicationID string,
	ownerID string,
	status string,
) error {

	return repository.UpdateApplicationStatus(
		applicationID,
		ownerID,
		status,
	)
}