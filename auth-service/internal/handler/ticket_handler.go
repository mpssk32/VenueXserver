package handler

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"
)

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