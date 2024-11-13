package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/thneutral/sdst/code/server/internal/database"
	"github.com/thneutral/sdst/code/server/internal/editorhub"
)

func HandleEditorHub(hub *editorhub.Hub) WSAuthenticatedHandler {
	return func(conn *websocket.Conn, user database.User) {
		type RequestModel struct {
			FileName string `json:"filename"`
		}
		var reqmodel RequestModel
		err := conn.ReadJSON(&reqmodel)
		if err != nil {
			wsWriteError(conn, websocket.CloseGoingAway, "Failed to recieve message")
			return
		}
		responseChannel := make(chan editorhub.TAddResponse)
		hub.AddRequest <- editorhub.TAddRequest{
			FileName: reqmodel.FileName,
			Conn:     conn,
			Response: responseChannel,
		}
		response := <-responseChannel
		if response.Error != nil {
			wsWriteError(conn, websocket.CloseGoingAway, "Failed to open file")
			return
		}
		defer func() {
			hub.DeleteRequest <- editorhub.TDeleteRequest{
				FileName: reqmodel.FileName,
				Conn:     conn,
			}
		}()
		type ResponseModel struct {
			Content []string `json:"content"`
		}
		var respmodel ResponseModel
		for _, data := range response.Content {
			respmodel.Content = append(respmodel.Content, string(data))
		}
		err = conn.WriteJSON(respmodel)
		if err != nil {
			fmt.Println("Failed to write message")
			return
		}
		for {
			messageType, data, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Failed to read message")
				return
			}
			if messageType == websocket.CloseMessage {
				return
			}
			var diff struct {
				Difference []editorhub.TDifference `json:"diff"`
			}
			_ = json.Unmarshal(data, &diff)
			hub.WriteRequest <- editorhub.TWriteRequest{
				FileName:   reqmodel.FileName,
				Difference: diff.Difference,
				Conn:       conn,
			}
			var resp struct {
				Ack string `json:"ack"`
			}
			resp.Ack = "Acknowledged"
			json, _ := json.Marshal(resp)
			conn.WriteMessage(websocket.TextMessage, json)
		}
	}
}
