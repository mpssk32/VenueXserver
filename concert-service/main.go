//	@title			VenueX Concert API
//	@version		1.0
//	@description	Concert Service
//	@host			localhost:8082
//	@BasePath		/
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

package main

import (
	_ "concert-service/docs"

	"concert-service/internal/config"
	"venuex/shared/logger"
	"concert-service/internal/routes"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	logger.Init()

	logger.Log.Info(
		"starting concert-service",
	)

	config.ConnectDB()

	logger.Log.Info(
		"PostgreSQL connected",
	)

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(
		httpSwagger.WrapHandler,
	)

	routes.RegisterRoutes(r)

	logger.Log.Infow(
		"concert-service started",
		"port",
		":8082",
	)

	err := http.ListenAndServe(
		":8082",
		r,
	)

	if err != nil {

		logger.Log.Fatalw(
			"failed to start concert-service",
			"error",
			err,
		)
	}
}