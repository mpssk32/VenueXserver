package routes

import (
	"net/http"

	"concert-service/internal/handler"
	"concert-service/internal/middleware"
	sharedjwt "venuex/shared/jwt"

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

	// создание площадок для владельцев площадок
	r.Handle(
	"/venues",
	sharedjwt.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.CreateVenue),
		),
	),
	).Methods("POST")

// создание концертов для владельцев площадок
	r.Handle(
	"/concerts",
	sharedjwt.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.CreateConcert),
		),
	),
	).Methods("POST")

	r.HandleFunc("/concerts", handler.GetConcerts).Methods("GET")

// создание заявок на проведение концерта для артистов
	r.Handle(
	"/applications",
	sharedjwt.JWTMiddleware(
		middleware.RoleMiddleware("artist")(
			http.HandlerFunc(handler.CreateApplication),
		),
	),
	).Methods("POST")


// просмотр заявок на проведение концерта для владельцев площадок
	r.Handle(
	"/applications",
	sharedjwt.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.GetVenueApplications),
		),
	),
	).Methods("GET")


// принятие и отклонение заявок на проведение концерта

	r.Handle(
	"/applications/{id}/approve",
	sharedjwt.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.ApproveApplication),
		),
	),
	).Methods("PATCH")

	r.Handle(
	"/applications/{id}/reject",
	sharedjwt.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.RejectApplication),
		),
	),
	).Methods("PATCH")

	// создание событий для владельцев площадок bearer token
	r.Handle(
		"/events",
		sharedjwt.JWTMiddleware(
			middleware.RoleMiddleware("venue_admin")(
				http.HandlerFunc(handler.CreateEvent),
			),
		),
	).Methods("POST")


	// получение всех событий для всех пользователей 
	r.HandleFunc(
		"/events",
		handler.GetAllEvents,
	).Methods("GET")

	// создание билетов bearer token 
	r.Handle(
		"/tickets",
		sharedjwt.JWTMiddleware(
			http.HandlerFunc(handler.CreateTicket),
		),
	).Methods("POST")

	
// получение билетов пользователя bearer token 

	r.Handle(
		"/my-tickets",
		sharedjwt.JWTMiddleware(
			http.HandlerFunc(handler.GetUserTickets),
		),
	).Methods("GET")


	// получение события по id для всех пользователей 
	r.HandleFunc(
		"/events/{id}",
		handler.GetEventByID,
	).Methods("GET")

	// удаление события  Bearer VENUE_ADMIN_TOKEN 
	r.Handle(
		"/events/{id}",
		sharedjwt.JWTMiddleware(
			middleware.RoleMiddleware("venue_admin")(
				http.HandlerFunc(handler.DeleteEvent),
			),
		),
	).Methods("DELETE")
	
	// обновление события Bearer VENUE_ADMIN_TOKEN 
	r.Handle(
		"/events/{id}",
		sharedjwt.JWTMiddleware(
			middleware.RoleMiddleware("venue_admin")(
				http.HandlerFunc(handler.UpdateEvent),
			),
		),
	).Methods("PUT")

	// поиск событий по названию для всех пользователей  http://localhost:8080/events/search?title=techno
	r.HandleFunc( 
		"/events/search",
		handler.SearchEvents,
	).Methods("GET")




	r.Handle(
		"/events",
		sharedjwt.JWTMiddleware(
			middleware.RoleMiddleware(
				"venue_admin",
			)(
			http.HandlerFunc(handler.CreateEvent),
			),
		),
	).Methods("POST")
}
