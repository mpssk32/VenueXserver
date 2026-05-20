package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func GetMessages(user1, user2 string) ([]models.Message, error) {
	return repository.GetMessages(user1, user2)
}