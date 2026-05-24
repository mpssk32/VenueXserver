package repository

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"context"
)

func CreateConcert(concert models.Concert) error {
	query := `
		INSERT INTO concerts (
			title,
			description,
			genre,
			concert_date,
			ticket_price,
			venue_id,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		concert.Title,
		concert.Description,
		concert.Genre,
		concert.ConcertDate,
		concert.TicketPrice,
		concert.VenueID,
		concert.CreatedBy,
	)

	return err
}

func GetConcerts() ([]models.Concert, error) {
	query := `
		SELECT
			id,
			title,
			description,
			genre,
			concert_date,
			ticket_price,
			venue_id,
			created_by,
			created_at
		FROM concerts
		ORDER BY concert_date ASC
	`

	rows, err := config.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var concerts []models.Concert

	for rows.Next() {
		var concert models.Concert

		err := rows.Scan(
			&concert.ID,
			&concert.Title,
			&concert.Description,
			&concert.Genre,
			&concert.ConcertDate,
			&concert.TicketPrice,
			&concert.VenueID,
			&concert.CreatedBy,
			&concert.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		concerts = append(concerts, concert)
	}

	return concerts, nil
}