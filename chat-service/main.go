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
	_ "chat-service/docs"

	"chat-service/internal/config"
	"chat-service/internal/handler"
	"venuex/shared/logger"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	logger.Init()

	logger.Log.Info(
		"starting chat-service",
	)

	config.ConnectDB()

	logger.Log.Info(
		"PostgreSQL connected",
	)

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

	logger.Log.Infow(
		"chat-service started",
		"port",
		":8081",
	)

	err := http.ListenAndServe(
		":8081",
		r,
	)

	if err != nil {

		logger.Log.Fatalw(
			"failed to start chat-service",
			"error",
			err,
		)
	}
}