package repository

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"context"
)

func CreateUser(user models.User) error {
	query := `
		INSERT INTO users (username, email, password, role)
		VALUES ($1, $2, $3, $4)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
	)

	return err
}

func GetUserByEmail(email string) (models.User, error) {
	query := `
		SELECT id, username, email, password, role, created_at
		FROM users
		WHERE email=$1
	`

	var user models.User

	err := config.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	return user, err
}

func GetUserByID(id string) (models.User, error) {
	query := `
		SELECT id, username, email, password, role, created_at
		FROM users
		WHERE id=$1
	`

	var user models.User
	err := config.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)


	return user, err
}