package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func CreateTicket(ticket models.Ticket) error {
	return repository.CreateTicket(ticket)
}

func GetUserTickets(userID string) ([]models.Ticket, error) {
	return repository.GetUserTickets(userID)
}