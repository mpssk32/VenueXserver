package repository

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"context"
)

func CreateAuditLog(log models.AuditLog) error {

	query := `
		INSERT INTO audit_logs (
			user_id,
			action
		)
		VALUES ($1, $2)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		log.UserID,
		log.Action,
	)

	return err
}