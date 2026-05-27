//	@title			VenueX Auth API
//	@version		1.0
//	@description	Authentication Service
//	@host			localhost:8080
//	@BasePath		/
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

package main

import (
	_ "auth-service/docs"

	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
	"net/http"

	"venuex/shared/logger"
	sharedjwt "venuex/shared/jwt"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	logger.Init()

	logger.Log.Info(
		"starting auth-service",
	)

	config.ConnectDB()

	logger.Log.Info(
		"PostgreSQL connected",
	)

	config.ConnectRabbitMQ()

	logger.Log.Info(
		"RabbitMQ connected",
	)

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(
		httpSwagger.WrapHandler,
	)

	r.HandleFunc(
		"/register",
		handler.Register,
	).Methods("POST")

	r.HandleFunc(
		"/login",
		handler.Login,
	).Methods("POST")

	r.HandleFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {

			w.Write(
				[]byte("auth service ok"),
			)
		},
	)

	// admin routes
	r.Handle(
		"/admin",
		sharedjwt.JWTMiddleware(
			middleware.RoleMiddleware("admin")(
				http.HandlerFunc(
					handler.AdminPanel,
				),
			),
		),
	).Methods("GET")

	// current user
	r.Handle(
		"/me",
		sharedjwt.JWTMiddleware(
			http.HandlerFunc(
				handler.GetMe,
			),
		),
	).Methods("GET")

	// super admin routes
	r.Handle(
		"/admin/users",
		sharedjwt.JWTMiddleware(
			middleware.RoleMiddleware("super_admin")(
				http.HandlerFunc(
					handler.GetAllUsers,
				),
			),
		),
	).Methods("GET")

	r.Handle(
		"/admin/users/{id}",
		sharedjwt.JWTMiddleware(
			middleware.RoleMiddleware("super_admin")(
				http.HandlerFunc(
					handler.DeleteUser,
				),
			),
		),
	).Methods("DELETE")

	logger.Log.Infow(
		"auth-service started",
		"port",
		":8080",
	)

	err := http.ListenAndServe(
		":8080",
		r,
	)

	if err != nil {

		logger.Log.Fatalw(
			"failed to start auth-service",
			"error",
			err,
		)
	}
}