package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thneutral/sdst/code/server/internal/database"
)

type WSAuthenticatedHandler func(conn *websocket.Conn, user database.User)

func WSGateway(upgrader websocket.Upgrader, queries *database.Queries, handler WSAuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			writeError(w, "Failed to upgrade http connection to ws", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		type TAuthMessage struct {
			Token string `json:"token"`
		}
		messageChan := make(chan TAuthMessage)

		go func() {
			var authenticationMessage TAuthMessage
			err = conn.ReadJSON(&authenticationMessage)
			messageChan <- authenticationMessage
		}()

		timer := time.NewTimer(1 * time.Second)
		select {
		case <-timer.C:
			{
				wsWriteError(conn, websocket.CloseGoingAway, "Failed to recieve authentication message")
				return
			}
		case message := <-messageChan:
			{
				if err != nil {
					wsWriteError(conn, websocket.CloseGoingAway, "Failed to process message")
					return
				}
				token, err := uuid.Parse(message.Token)
				if err != nil {
					wsWriteError(conn, websocket.CloseGoingAway, "Invalid token format")
					return
				}
				user, err := queries.GetUserByToken(r.Context(), token)
				if err != nil {
					wsWriteError(conn, websocket.CloseGoingAway, "User with this token was not found")
					return
				}
				conn.WriteJSON(struct {
					Message string `json:"message"`
				}{
					Message: "Successfully logged in",
				})
				handler(conn, user)
			}
		}
	}
}

func wsWriteError(conn *websocket.Conn, messageType int, message string) {
	var payload struct {
		ErrorMessage string `json:"error_message"`
	}
	payload.ErrorMessage = message
	json, _ := json.Marshal(payload)
	conn.WriteMessage(messageType, json)
}
