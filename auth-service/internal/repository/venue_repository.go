package repository

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
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