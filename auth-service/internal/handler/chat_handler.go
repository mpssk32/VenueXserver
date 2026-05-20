package handler

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
}

var clients = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "user_id required", http.StatusBadRequest)
		return
	}


	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	clients[userID] = conn

	log.Println("Client connected")

	for {

		_, payload, err := conn.ReadMessage()

		if err != nil {

			delete(clients, userID)

			log.Println("Client disconnected")

			break
		}

		var chatMessage ChatMessage

		err = json.Unmarshal(payload, &chatMessage)

		if err != nil {
			log.Println(err)
			continue
		}

		message := models.Message{
			SenderID:   chatMessage.SenderID,
			ReceiverID: chatMessage.ReceiverID,
			Content:    chatMessage.Content,
		}

		err = repository.SaveMessage(message)

		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf(
			"Message from %s to %s: %s\n",
			chatMessage.SenderID,
			chatMessage.ReceiverID,
			chatMessage.Content,
		)

		receiverConn, ok := clients[chatMessage.ReceiverID]

		if ok {

			err := receiverConn.WriteMessage(
				websocket.TextMessage,
				payload,
			)

			if err != nil {

				receiverConn.Close()

				delete(clients, chatMessage.ReceiverID)
			}
		}
	}
}