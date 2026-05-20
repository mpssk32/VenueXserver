package handler

import (
	"auth-service/internal/service"
	"encoding/json"
	"net/http"
)

func GetMe(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(string)

	user, err := service.GetUserByID(userID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}