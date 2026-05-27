package handler

import (
	"chat-service/internal/models"
	"chat-service/internal/service"
	"encoding/json"
	"net/http"

	sharedjwt "venuex/shared/jwt"
	"venuex/shared/logger"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

// подключение в постмане через jwt токен, который выдаёт auth-service при логине, и отправка в теле сообщения вида:
//
//	{
//		"sender_id": "id_пользователя_из_токена",
//		"receiver_id": "id_получателя",
//		"content": "текст сообщения"
//	}

var clients = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ChatHandler godoc
//
//	@Summary		WebSocket chat
//	@Description	Real-time websocket chat connection
//	@Tags			chat
//	@Produce		json
//	@Param			token	query		string	true	"JWT token"
//	@Success		101		{string}	string	"Switching Protocols"
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Router			/ws [get]
func ChatHandler(w http.ResponseWriter, r *http.Request) {

	tokenString := r.URL.Query().Get("token")

	if tokenString == "" {

		logger.Log.Warn(
			"websocket connection without token",
		)

		http.Error(
			w,
			"token required",
			http.StatusBadRequest,
		)

		return
	}

	token, err := sharedjwt.ValidateJWT(
		tokenString,
	)

	if err != nil || !token.Valid {

		logger.Log.Warnw(
			"invalid jwt token",
			"error",
			err,
		)

		http.Error(
			w,
			"Invalid token",
			http.StatusUnauthorized,
		)

		return
	}

	claims := token.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)

	conn, err := upgrader.Upgrade(
		w,
		r,
		nil,
	)

	if err != nil {

		logger.Log.Errorw(
			"websocket upgrade failed",
			"error",
			err,
		)

		return
	}

	defer conn.Close()

	clients[userID] = conn

	logger.Log.Infow(
		"client connected",
		"user_id",
		userID,
	)

	for {

		_, payload, err := conn.ReadMessage()

		if err != nil {

			delete(
				clients,
				userID,
			)

			logger.Log.Infow(
				"client disconnected",
				"user_id",
				userID,
			)

			break
		}

		var chatMessage models.Message

		err = json.Unmarshal(
			payload,
			&chatMessage,
		)

		if err != nil {

			logger.Log.Warnw(
				"invalid websocket payload",
				"error",
				err,
			)

			continue
		}

		if chatMessage.SenderID != userID {

			logger.Log.Warnw(
				"sender_id mismatch",
				"user_id",
				userID,
				"sender_id",
				chatMessage.SenderID,
			)

			http.Error(
				w,
				"sender_id mismatch",
				http.StatusUnauthorized,
			)

			return
		}

		message := models.Message{
			SenderID:   chatMessage.SenderID,
			ReceiverID: chatMessage.ReceiverID,
			Content:    chatMessage.Content,
		}

		err = service.MessageService.SaveMessage(
			message,
		)

		if err != nil {

			logger.Log.Errorw(
				"failed to save message",
				"error",
				err,
			)

			continue
		}

		receiverConn, ok := clients[
			chatMessage.ReceiverID]

		if ok {

			err := receiverConn.WriteMessage(
				websocket.TextMessage,
				payload,
			)

			if err != nil {

				logger.Log.Errorw(
					"failed to send websocket message",
					"error",
					err,
				)

				receiverConn.Close()

				delete(
					clients,
					chatMessage.ReceiverID,
				)
			}
		}
	}
}