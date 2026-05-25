package handler

import (
	"concert-service/internal/models"
	"concert-service/internal/service"
	"encoding/json"
	"net/http"
)
// CreateConcert godoc
//
//	@Summary		Create concert
//	@Description	Create new concert
//	@Tags			concerts
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		plain
//	@Param			request	body		models.Concert	true	"Concert body"
//	@Success		201		{string}	string	"Concert created"
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Failure		500		{string}	string
//	@Router			/concerts [post]
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

// GetConcerts godoc
//
//	@Summary		Get concerts
//	@Description	Get all concerts
//	@Tags			concerts
//	@Produce		json
//	@Success		200	{array}		models.Concert
//	@Failure		500	{string}	string
//	@Router			/concerts [get]
func GetConcerts(w http.ResponseWriter, r *http.Request) {
	concerts, err := service.GetConcerts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(concerts)
}