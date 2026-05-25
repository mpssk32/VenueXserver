package handler

import (
	"concert-service/internal/models"
	"concert-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


// CreateEvent godoc
//
//	@Summary		Create event
//	@Description	Create new concert event
//	@Tags			events
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		plain
//	@Param			request	body		models.Event	true	"Event body"
//	@Success		201		{string}	string	"Event created"
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Failure		500		{string}	string
//	@Router			/events [post]
func CreateEvent(w http.ResponseWriter, r *http.Request) {

	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	if event.Title == "" ||
		event.VenueID == "" ||
		event.ArtistID == "" {

		http.Error(
			w,
			"All fields are required",
			http.StatusBadRequest,
		)

		return
	}

	err = service.CreateEvent(event)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Event created"))
}


// GetAllEvents godoc
//
//	@Summary		Get all events
//	@Description	Get all concert events
//	@Tags			events
//	@Produce		json
//	@Success		200	{array}		models.Event
//	@Failure		500	{string}	string
//	@Router			/events [get]
func GetAllEvents(w http.ResponseWriter, r *http.Request) {

	events, err := service.GetAllEvents()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(events)
}


// GetEventByID godoc
//
//	@Summary		Get event by ID
//	@Description	Get single event
//	@Tags			events
//	@Produce		json
//	@Param			id	path		int	true	"Event ID"
//	@Success		200	{object}	models.Event
//	@Failure		404	{string}	string
//	@Router			/events/{id} [get]
func GetEventByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	eventID := vars["id"]

	event, err := service.GetEventByID(eventID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(event)
}


// DeleteEvent godoc
//
//	@Summary		Delete event
//	@Description	Delete event by ID
//	@Tags			events
//	@Security		BearerAuth
//	@Produce		plain
//	@Param			id	path		int	true	"Event ID"
//	@Success		200	{string}	string	"Event deleted"
//	@Failure		401	{string}	string
//	@Failure		403	{string}	string
//	@Failure		404	{string}	string
//	@Failure		500	{string}	string
//	@Router			/events/{id} [delete]
func DeleteEvent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	eventID := vars["id"]

	event, err := service.GetEventByID(eventID)

	if err != nil {

		http.Error(
			w,
			"Event not found",
			http.StatusNotFound,
		)

		return
	}

	userID := r.Context().Value("user_id").(string)

	if event.VenueID != userID {

		http.Error(
			w,
			"Forbidden",
			http.StatusForbidden,
		)

		return
	}

	err = service.DeleteEvent(eventID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Write([]byte("Event deleted"))
}


// UpdateEvent godoc
//
//	@Summary		Update event
//	@Description	Update event by ID
//	@Tags			events
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		plain
//	@Param			id		path		int				true	"Event ID"
//	@Param			request	body		models.Event	true	"Updated event"
//	@Success		200		{string}	string	"Event updated"
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Failure		403		{string}	string
//	@Failure		404		{string}	string
//	@Failure		500		{string}	string
//	@Router			/events/{id} [put]
func UpdateEvent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	eventID := vars["id"]

	existingEvent, err := service.GetEventByID(eventID)

	if err != nil {

		http.Error(
			w,
			"Event not found",
			http.StatusNotFound,
		)

		return
	}

	userID := r.Context().Value("user_id").(string)

	if existingEvent.VenueID != userID {

		http.Error(
			w,
			"Forbidden",
			http.StatusForbidden,
		)

		return
	}

	var event models.Event

	err = json.NewDecoder(r.Body).Decode(&event)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	err = service.UpdateEvent(eventID, event)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Write([]byte("Event updated"))
}


// SearchEvents godoc
//
//	@Summary		Search events
//	@Description	Search events by title
//	@Tags			events
//	@Produce		json
//	@Param			title	query		string	true	"Event title"
//	@Success		200		{array}	models.Event
//	@Failure		400		{string}	string
//	@Failure		500		{string}	string
//	@Router			/events/search [get]
func SearchEvents(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Query().Get("title")

	if title == "" {

		http.Error(
			w,
			"title query required",
			http.StatusBadRequest,
		)

		return
	}

	events, err := service.SearchEvents(title)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(events)
}