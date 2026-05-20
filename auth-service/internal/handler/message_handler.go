package handler

import (
	"auth-service/internal/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user1 := vars["user1"]
	user2 := vars["user2"]

	messages, err := repository.GetMessages(user1, user2)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(messages)
}