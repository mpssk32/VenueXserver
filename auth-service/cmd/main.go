// @title VenueX API
// @version 1.0
// @description Concert platform microservices API
// @host localhost:8000
// @BasePath /

package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
	"log"
	"net/http"

	_ "auth-service/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.ConnectDB()
	config.ConnectRabbitMQ()
	
	r := mux.NewRouter()
	
	r.PathPrefix("/swagger/").Handler(
		httpSwagger.WrapHandler,
	)

	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth service ok"))
	})

// админ панель для управления пользователями
	r.Handle("/admin",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("admin")(
			http.HandlerFunc(handler.AdminPanel),
		),
	),
	).Methods("GET")


// создание площадок для владельцев площадок
	r.Handle(
	"/venues",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.CreateVenue),
		),
	),
	).Methods("POST")

// создание концертов для владельцев площадок
	r.Handle(
	"/concerts",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.CreateConcert),
		),
	),
	).Methods("POST")

	r.HandleFunc("/concerts", handler.GetConcerts).Methods("GET")



// создание заявок на проведение концерта для артистов
	r.Handle(
	"/applications",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("artist")(
			http.HandlerFunc(handler.CreateApplication),
		),
	),
	).Methods("POST")


// просмотр заявок на проведение концерта для владельцев площадок
	r.Handle(
	"/applications",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.GetVenueApplications),
		),
	),
	).Methods("GET")

// принятие и отклонение заявок на проведение концерта

	r.Handle(
	"/applications/{id}/approve",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.ApproveApplication),
		),
	),
	).Methods("PATCH")

	r.Handle(
	"/applications/{id}/reject",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.RejectApplication),
		),
	),
	).Methods("PATCH")

	// создание событий для владельцев площадок bearer token
	r.Handle(
		"/events",
		middleware.JWTMiddleware(
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
		middleware.JWTMiddleware(
			http.HandlerFunc(handler.CreateTicket),
		),
	).Methods("POST")

	// получение билетов пользователя bearer token 
	r.Handle(
		"/my-tickets",
		middleware.JWTMiddleware(
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
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("venue_admin")(
				http.HandlerFunc(handler.DeleteEvent),
			),
		),
	).Methods("DELETE")
	
	// обновление события Bearer VENUE_ADMIN_TOKEN 
	r.Handle(
		"/events/{id}",
		middleware.JWTMiddleware(
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

	// получение информации о себе для всех пользователей Bearer TOKEN
	r.Handle(
		"/me",
		middleware.JWTMiddleware(
			http.HandlerFunc(handler.GetMe),
		),
	).Methods("GET")

// админ панель для супер админов для управления всеми пользователями и их ролями
	r.Handle(
		"/admin/users",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("super_admin")(
			http.HandlerFunc(handler.GetAllUsers),
			),
		),
	).Methods("GET")

	// удаление пользователя супер админом Bearer SUPER_ADMIN_TOKEN
	r.Handle(
		"/admin/users/{id}",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("super_admin")(
				http.HandlerFunc(handler.DeleteUser),
			),
		),
	).Methods("DELETE")


	log.Println("Auth service started on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}