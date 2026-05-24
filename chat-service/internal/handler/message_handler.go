package handler

import (
	"chat-service/internal/service"
	"encoding/json"
	"net/http"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {

	senderID := r.URL.Query().Get("sender_id")
	receiverID := r.URL.Query().Get("receiver_id")

	messages, err := service.GetMessages(
		senderID,
		receiverID,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(messages)
}