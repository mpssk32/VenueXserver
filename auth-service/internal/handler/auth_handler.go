package handler

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"
)

// Register godoc
//
//	@Summary		Register new user
//	@Description	Create new account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.RegisterRequest	true	"Register body"
//	@Success		201		{string}	string	"User created"
//	@Failure		400		{string}	string
//	@Failure		500		{string}	string
//	@Router			/register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" ||
		req.Email == "" ||
		req.Password == "" ||
		req.Role == "" {

		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	err = service.AuthService.Register(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Login with email and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.LoginRequest	true	"Login body"
//	@Success		200		{object}	models.AuthResponse
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Router			/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	token, err := service.AuthService.Login(req)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}

// AdminPanel godoc
//
//	@Summary		Admin panel
//	@Description	Access admin panel
//	@Tags			admin
//	@Security		BearerAuth
//	@Produce		plain
//	@Success		200	{string}	string	"Welcome admin"
//	@Failure		401	{string}	string
//	@Failure		403	{string}	string
//	@Router			/admin [get]
func AdminPanel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome admin"))
}