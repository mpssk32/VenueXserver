//	@title			VenueX Chat API
//	@version		1.0
//	@description	Chat Service
//	@host			localhost:8081
//	@BasePath		/
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
package main

import (
	"chat-service/internal/config"
	"chat-service/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "chat-service/docs"
)

func main() {

	config.ConnectDB()

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(
		httpSwagger.WrapHandler,
	)


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