package handler

import (
	"concert-service/internal/models"
	"concert-service/internal/service"
	"encoding/json"
	"net/http"
)

func CreateConcert(w http.ResponseWriter, r *http.Request) {
	var concert models.Concert

	err := json.NewDecoder(r.Body).Decode(&concert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if concert.Title == "" ||
		concert.Description == "" ||
		concert.Genre == "" ||
		concert.VenueID == "" ||
		concert.TicketPrice < 0 {

		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	concert.CreatedBy = r.Context().Value("user_id").(string)

	err = service.CreateConcert(concert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Concert created"))
}

func GetConcerts(w http.ResponseWriter, r *http.Request) {
	concerts, err := service.GetConcerts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(concerts)
}