package models

import "time"

type Ticket struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	EventID   string    `json:"event_id"`
	CreatedAt time.Time `json:"created_at"`
}