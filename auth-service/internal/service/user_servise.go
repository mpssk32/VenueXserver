package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func GetUserByID(userID string) (models.User, error) {
	return repository.GetUserByID(userID)
}

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

func DeleteUser(userID string) error {
	return repository.DeleteUser(userID)
}