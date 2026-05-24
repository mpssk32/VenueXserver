package main

import (
	"chat-service/internal/config"
	"chat-service/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDB()

	r := mux.NewRouter()

	// websocket chat
	r.HandleFunc(
		"/ws",
		handler.ChatHandler,
	)

	// http://localhost:8081/messages?sender_id=UUID1&receiver_id=UUID2
	r.HandleFunc(
		"/messages",
		handler.GetMessages,
	).Methods("GET")

	log.Println("чат запущен на :8081")

	log.Fatal(
		http.ListenAndServe(":8081", r),
	)
}