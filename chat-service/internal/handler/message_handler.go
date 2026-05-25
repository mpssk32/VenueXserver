package handler

import (
	_ "chat-service/internal/models"
	"chat-service/internal/service"
	"encoding/json"
	"net/http"
)


// GetMessages godoc
//
//	@Summary		Get messages
//	@Description	Get chat history between users
//	@Tags			messages
//	@Produce		json
//	@Param			sender_id		query		string	true	"Sender ID"
//	@Param			receiver_id	query		string	true	"Receiver ID"
//	@Success		200				{array}	models.Message
//	@Failure		500				{string}	string
//	@Router			/messages [get]
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