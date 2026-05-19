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
	r.Handle("/admin",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("admin")(
			http.HandlerFunc(handler.AdminPanel),
		),
	),
	).Methods("GET")

	r.Handle(
	"/venues",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.CreateVenue),
		),
	),
	).Methods("POST")

	r.Handle(
	"/concerts",
	middleware.JWTMiddleware(
		middleware.RoleMiddleware("venue_admin")(
			http.HandlerFunc(handler.CreateConcert),
		),
	),
	).Methods("POST")

	r.HandleFunc("/concerts", handler.GetConcerts).Methods("GET")
	
	log.Println("Auth service started on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}