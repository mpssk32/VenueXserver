package handler

import (
	"concert-service/internal/models"
	"concert-service/internal/service"
	"encoding/json"
	"net/http"
)


// CreateTicket godoc
//
//	@Summary		Book ticket
//	@Description	Book concert ticket
//	@Tags			tickets
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		plain
//	@Param			request	body		models.Ticket	true	"Ticket body"
//	@Success		201		{string}	string	"Ticket booked"
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Failure		500		{string}	string
//	@Router			/tickets [post]
func CreateTicket(w http.ResponseWriter, r *http.Request) {

	var ticket models.Ticket

	err := json.NewDecoder(r.Body).Decode(&ticket)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	if ticket.EventID == "" {

		http.Error(
			w,
			"All fields are required",
			http.StatusBadRequest,
		)

		return
	}
	userID := r.Context().Value("user_id").(string)

	ticket.UserID = userID

	err = service.CreateTicket(ticket)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Ticket booked"))
}

// GetUserTickets godoc
//
//	@Summary		Get user tickets
//	@Description	Get all booked tickets
//	@Tags			tickets
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID"
//	@Success		200			{array}	models.Ticket
//	@Failure		400			{string}	string
//	@Failure		500			{string}	string
//	@Router			/tickets [get]
func GetUserTickets(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	if userID == "" {

		http.Error(
			w,
			"user_id required",
			http.StatusBadRequest,
		)

		return
	}

	tickets, err := service.GetUserTickets(userID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tickets)
}