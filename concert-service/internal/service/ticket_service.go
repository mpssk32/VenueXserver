package service

import (
	"concert-service/internal/models"
	"concert-service/internal/repository"
)

func CreateTicket(ticket models.Ticket) error {
	return repository.CreateTicket(ticket)
}

func GetUserTickets(userID string) ([]models.Ticket, error) {
	return repository.GetUserTickets(userID)
}