package models

import "time"

type Application struct {
	ID        string    `json:"id"`
	ArtistID  string    `json:"artist_id"`
	VenueID   string    `json:"venue_id"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}