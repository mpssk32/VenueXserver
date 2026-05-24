// @title VenueX API
// @version 1.0
// @description Concert platform microservices API
// @host localhost:8000
// @BasePath /

package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
	"log"
	"net/http"

	_ "auth-service/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.ConnectDB()
	config.ConnectRabbitMQ()
	
	r := mux.NewRouter()
	
	r.PathPrefix("/swagger/").Handler(
		httpSwagger.WrapHandler,
	)

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

	// получение информации о себе для всех пользователей Bearer TOKEN
	r.Handle(
		"/me",
		middleware.JWTMiddleware(
			http.HandlerFunc(handler.GetMe),
		),
	).Methods("GET")

// админ панель для супер админов для управления всеми пользователями и их ролями
	r.Handle(
		"/admin/users",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("super_admin")(
			http.HandlerFunc(handler.GetAllUsers),
			),
		),
	).Methods("GET")

	// удаление пользователя супер админом Bearer SUPER_ADMIN_TOKEN
	r.Handle(
		"/admin/users/{id}",
		middleware.JWTMiddleware(
			middleware.RoleMiddleware("super_admin")(
				http.HandlerFunc(handler.DeleteUser),
			),
		),
	).Methods("DELETE")


	log.Println("Auth service started on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}