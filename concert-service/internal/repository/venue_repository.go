package repository

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"context"
)

func CreateVenue(venue models.Venue) error {
	query := `
		INSERT INTO venues (
			name,
			city,
			description,
			equipment,
			capacity,
			owner_id
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		venue.Name,
		venue.City,
		venue.Description,
		venue.Equipment,
		venue.Capacity,
		venue.OwnerID,
	)

	return err
}

func GetVenues() ([]models.Venue, error) {
	query := `
		SELECT
			id,
			name,
			city,
			description,
			equipment,
			capacity,
			owner_id
		FROM venues
	`
	rows, err := config.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var venues []models.Venue
	for rows.Next() {
		var v models.Venue
		err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.City,
			&v.Description,
			&v.Equipment,
			&v.Capacity,
			&v.OwnerID,
		)
		if err != nil {
			return nil, err
		}
		venues = append(venues, v)
	}

	return venues, nil
}