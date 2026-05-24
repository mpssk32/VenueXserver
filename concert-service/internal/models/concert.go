package models

import "time"

type Concert struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Genre        string    `json:"genre"`
	ConcertDate  time.Time `json:"concert_date"`
	TicketPrice  int       `json:"ticket_price"`
	VenueID      string    `json:"venue_id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
}