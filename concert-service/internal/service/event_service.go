package service

import (
	"concert-service/internal/models"
	"concert-service/internal/repository"
)

func CreateEvent(event models.Event) error {
	return repository.CreateEvent(event)
}

func GetAllEvents() ([]models.Event, error) {
	return repository.GetAllEvents()
}

func GetEventByID(eventID string) (models.Event, error) {
	return repository.GetEventByID(eventID)
}

func DeleteEvent(eventID string) error {
	return repository.DeleteEvent(eventID)
}

func UpdateEvent(eventID string, event models.Event) error {
	return repository.UpdateEvent(eventID, event)
}

func SearchEvents(title string) ([]models.Event, error) {
	return repository.SearchEvents(title)
}