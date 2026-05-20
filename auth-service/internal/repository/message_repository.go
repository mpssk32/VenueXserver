package repository

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"context"
)


func SaveMessage(message models.Message) error {

	query := `
		INSERT INTO messages (
			sender_id,
			receiver_id,
			content
		)
		VALUES ($1, $2, $3)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		message.SenderID,
		message.ReceiverID,
		message.Content,
	)

	return err
}

func GetMessages(user1, user2 string) ([]models.Message, error) {

	query := `
		SELECT
			id,
			sender_id,
			receiver_id,
			content,
			created_at
		FROM messages
		WHERE
			(sender_id = $1 AND receiver_id = $2)
			OR
			(sender_id = $2 AND receiver_id = $1)
		ORDER BY created_at ASC
	`

	rows, err := config.DB.Query(
		context.Background(),
		query,
		user1,
		user2,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []models.Message

	for rows.Next() {

		var message models.Message

		err := rows.Scan(
			&message.ID,
			&message.SenderID,
			&message.ReceiverID,
			&message.Content,
			&message.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}


