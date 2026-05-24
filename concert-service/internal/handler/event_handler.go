package handler

import (
	"concert-service/internal/models"
	"concert-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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