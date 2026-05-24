package repository

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"context"
)

func CreateTicket(ticket models.Ticket) error {

	query := `
		INSERT INTO tickets (
			user_id,
			event_id
		)
		VALUES ($1, $2)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		ticket.UserID,
		ticket.EventID,
	)

	return err
}

func GetUserTickets(userID string) ([]models.Ticket, error) {

	query := `
		SELECT
			id,
			user_id,
			event_id,
			created_at
		FROM tickets
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := config.DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tickets []models.Ticket

	for rows.Next() {

		var ticket models.Ticket

		err := rows.Scan(
			&ticket.ID,
			&ticket.UserID,
			&ticket.EventID,
			&ticket.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}