package repository

import (
	"concert-service/internal/config"
	"concert-service/internal/models"
	"context"
)

func CreateApplication(app models.Application) error {
	query := `
		INSERT INTO applications (
			artist_id,
			venue_id,
			message
		)
		VALUES ($1, $2, $3)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		app.ArtistID,
		app.VenueID,
		app.Message,
	)

	return err
}

func GetVenueApplications(ownerID string) ([]models.Application, error) {

	query := `
		SELECT
			applications.id,
			applications.artist_id,
			applications.venue_id,
			applications.message,
			applications.status,
			applications.created_at
		FROM applications
		JOIN venues
			ON applications.venue_id = venues.id
		WHERE venues.owner_id = $1
		ORDER BY applications.created_at DESC
	`

	rows, err := config.DB.Query(
		context.Background(),
		query,
		ownerID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var applications []models.Application

	for rows.Next() {

		var app models.Application

		err := rows.Scan(
			&app.ID,
			&app.ArtistID,
			&app.VenueID,
			&app.Message,
			&app.Status,
			&app.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		applications = append(applications, app)
	}

	return applications, nil
}


func UpdateApplicationStatus(
	applicationID string,
	ownerID string,
	status string,
) error {

	query := `
		UPDATE applications
		SET status = $1
		FROM venues
		WHERE applications.venue_id = venues.id
			AND applications.id = $2
			AND venues.owner_id = $3
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		status,
		applicationID,
		ownerID,
	)

	return err
}


func DeleteApplication(applicationID string) error {

	query := `
		DELETE FROM applications
		WHERE id = $1
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		applicationID,
	)

	return err
}