package handler

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"
)

func CreateVenue(w http.ResponseWriter, r *http.Request) {
	var venue models.Venue

	err := json.NewDecoder(r.Body).Decode(&venue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if venue.Name == "" ||
		venue.City == "" ||
		venue.Description == "" ||
		venue.Equipment == "" ||
		venue.Capacity == 0 {

		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	venue.OwnerID = r.Context().Value("user_id").(string)

	err = service.CreateVenue(venue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Venue created"))
}