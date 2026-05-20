package main

import (
	"log"
	"net/http"
	"auth-service/internal/config"
	"github.com/gorilla/mux"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
)

func main() {
	config.ConnectDB()
	
	r := mux.NewRouter()
	
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



	log.Println("Auth service started on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}