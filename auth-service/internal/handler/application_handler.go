package handler

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateApplication(w http.ResponseWriter, r *http.Request) {

	var app models.Application

	err := json.NewDecoder(r.Body).Decode(&app)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	if app.VenueID == "" || app.Message == "" {

		http.Error(
			w,
			"All fields are required",
			http.StatusBadRequest,
		)

		return
	}

	app.ArtistID = r.Context().Value("user_id").(string)

	err = service.CreateApplication(app)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Application created"))
}

func GetVenueApplications(w http.ResponseWriter, r *http.Request) {

	ownerID := r.Context().Value("user_id").(string)

	applications, err := service.GetVenueApplications(ownerID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(applications)
}

func ApproveApplication(w http.ResponseWriter, r *http.Request) {

	applicationID := mux.Vars(r)["id"]

	ownerID := r.Context().Value("user_id").(string)

	err := service.UpdateApplicationStatus(
		applicationID,
		ownerID,
		"approved",
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	err = service.PublishNotification(
		"Application approved",
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	err = service.DeleteApplication(applicationID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}
	
	err = service.DeleteApplication(applicationID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}


	w.Write([]byte("Application approved"))
}

func RejectApplication(w http.ResponseWriter, r *http.Request) {

	applicationID := mux.Vars(r)["id"]

	ownerID := r.Context().Value("user_id").(string)

	err := service.UpdateApplicationStatus(
		applicationID,
		ownerID,
		"rejected",
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	err = service.PublishNotification(
		"Application rejected",
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	err = service.DeleteApplication(applicationID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Write([]byte("Application rejected"))
}