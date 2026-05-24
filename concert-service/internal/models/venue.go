package models

import "time"

type Venue struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Description string    `json:"description"`
	Equipment   string    `json:"equipment"`
	Capacity    int       `json:"capacity"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
}