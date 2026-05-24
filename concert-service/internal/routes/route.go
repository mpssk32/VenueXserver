package routes

import (
	"net/http"

	"concert-service/internal/handler"
	"concert-service/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	r.HandleFunc(
		"/events",
		handler.GetAllEvents,
	).Methods("GET")

	r.HandleFunc(
		"/events/search",
		handler.SearchEvents,
	).Methods("GET")

	r.HandleFunc(
		"/events/{id}",
		handler.GetEventByID,
	).Methods("GET")

	r.Handle(
		"/events",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware(
				"venue_admin",
			)(
			http.HandlerFunc(handler.CreateEvent),
			),
		),
	).Methods("POST")
}
