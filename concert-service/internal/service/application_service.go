package service

import (
	"concert-service/internal/models"
	"concert-service/internal/repository"
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

func DeleteApplication(applicationID string) error {
	return repository.DeleteApplication(applicationID)
}