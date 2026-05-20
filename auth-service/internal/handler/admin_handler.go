package handler

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := service.GetAllUsers()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	userID := vars["id"]

	err := service.DeleteUser(userID)

if err != nil {

	http.Error(
		w,
		err.Error(),
		http.StatusInternalServerError,
	)

	return
}

	adminID := r.Context().Value("user_id").(string)

	service.CreateAuditLog(models.AuditLog{
		UserID: adminID,
		Action: "Deleted user " + userID,
	})

	w.Write([]byte("User deleted"))
}