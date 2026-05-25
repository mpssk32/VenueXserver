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
	"concert-service/internal/config"
	"concert-service/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "concert-service/docs"
)

func main() {

	config.ConnectDB()

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(
		httpSwagger.WrapHandler,
	)
	
	routes.RegisterRoutes(r)

	log.Println("concert service started on :8082")

	log.Fatal(
		http.ListenAndServe(":8082", r),
	)
}