package handler

import (
	"concert-service/internal/models"
	"concert-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateApplication godoc
//
//	@Summary		Create application
//	@Description	Artist sends application to venue
//	@Tags			applications
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		plain
//	@Param			request	body		models.Application	true	"Application body"
//	@Success		201		{string}	string	"Application created"
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Failure		500		{string}	string
//	@Router			/applications [post]
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


// GetVenueApplications godoc
//
//	@Summary		Get venue applications
//	@Description	Get applications for venue owner
//	@Tags			applications
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}		models.Application
//	@Failure		401	{string}	string
//	@Failure		500	{string}	string
//	@Router			/applications [get]
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


// ApproveApplication godoc
//
//	@Summary		Approve application
//	@Description	Approve artist application
//	@Tags			applications
//	@Security		BearerAuth
//	@Produce		plain
//	@Param			id	path		int	true	"Application ID"
//	@Success		200	{string}	string	"Application approved"
//	@Failure		401	{string}	string
//	@Failure		500	{string}	string
//	@Router			/applications/{id}/approve [put]
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

	// err = service.PublishNotification(
	// 	"Application approved",
	// )

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


// RejectApplication godoc
//
//	@Summary		Reject application
//	@Description	Reject artist application
//	@Tags			applications
//	@Security		BearerAuth
//	@Produce		plain
//	@Param			id	path		int	true	"Application ID"
//	@Success		200	{string}	string	"Application rejected"
//	@Failure		401	{string}	string
//	@Failure		500	{string}	string
//	@Router			/applications/{id}/reject [put]
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

	// err = service.PublishNotification(
	// 	"Application rejected",
	// )

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