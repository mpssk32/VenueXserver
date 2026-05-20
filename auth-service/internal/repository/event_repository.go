package repository

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"context"
)

func CreateEvent(event models.Event) error {

	query := `
		INSERT INTO events (
			title,
			description,
			venue_id,
			artist_id,
			event_date
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		event.Title,
		event.Description,
		event.VenueID,
		event.ArtistID,
		event.EventDate,
	)

	return err
}

func GetAllEvents() ([]models.Event, error) {

	query := `
		SELECT
			id,
			title,
			description,
			venue_id,
			artist_id,
			event_date,
			created_at
		FROM events
		ORDER BY event_date ASC
	`

	rows, err := config.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []models.Event

	for rows.Next() {

		var event models.Event

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.VenueID,
			&event.ArtistID,
			&event.EventDate,
			&event.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(eventID string) (models.Event, error) {

	query := `
		SELECT
			id,
			title,
			description,
			venue_id,
			artist_id,
			event_date,
			created_at
		FROM events
		WHERE id = $1
	`

	var event models.Event

	err := config.DB.QueryRow(
		context.Background(),
		query,
		eventID,
	).Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.VenueID,
		&event.ArtistID,
		&event.EventDate,
		&event.CreatedAt,
	)

	return event, err
}

func DeleteEvent(eventID string) error {

	query := `
		DELETE FROM events
		WHERE id = $1
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		eventID,
	)

	return err
}

func UpdateEvent(eventID string, event models.Event) error {

	query := `
		UPDATE events
		SET
			title = $1,
			description = $2,
			venue_id = $3,
			artist_id = $4,
			event_date = $5
		WHERE id = $6
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		event.Title,
		event.Description,
		event.VenueID,
		event.ArtistID,
		event.EventDate,
		eventID,
	)

	return err
}

func SearchEvents(title string) ([]models.Event, error) {

	query := `
		SELECT
			id,
			title,
			description,
			venue_id,
			artist_id,
			event_date,
			created_at
		FROM events
		WHERE LOWER(title) LIKE LOWER($1)
		ORDER BY event_date ASC
	`

	rows, err := config.DB.Query(
		context.Background(),
		query,
		"%"+title+"%",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []models.Event

	for rows.Next() {

		var event models.Event

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.VenueID,
			&event.ArtistID,
			&event.EventDate,
			&event.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}