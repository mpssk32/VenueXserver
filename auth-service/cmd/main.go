package main

import (
	"log"
	"net/http"
	"auth-service/internal/config"
	"github.com/gorilla/mux"
	"auth-service/internal/handler"
)

func main() {
	config.ConnectDB()
	
	r := mux.NewRouter()
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth service ok"))
	})

	log.Println("Auth service started on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}