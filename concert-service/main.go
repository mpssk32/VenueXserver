package main

import (
	"concert-service/internal/config"
	"concert-service/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDB()

	r := mux.NewRouter()

	routes.RegisterRoutes(r)

	log.Println("Event service started on :8082")

	log.Fatal(
		http.ListenAndServe(":8082", r),
	)
}