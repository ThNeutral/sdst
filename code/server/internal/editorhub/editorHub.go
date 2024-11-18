package editorhub

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type File struct {
	FileName string
	Content  string
	UsedBy   map[*websocket.Conn]int
}

type TDifference struct {
	Index int    `json:"index"`
	Data  string `json:"data"`
}

type TLockRequest struct {
	FileName   string
	Conn       *websocket.Conn
	LockedLine int
	By         string
}

type TWriteRequest struct {
	FileName string
	Conn     *websocket.Conn
	Data     string
}

type TAddResponse struct {
	Error   error
	Content string
}

type TAddRequest struct {
	FileName string
	Conn     *websocket.Conn
	Response chan TAddResponse
}

type TDeleteRequest struct {
	FileName string
	Conn     *websocket.Conn
}

type Hub struct {
	files         map[string]*File
	WriteRequest  chan TWriteRequest
	AddRequest    chan TAddRequest
	DeleteRequest chan TDeleteRequest
	LockRequest   chan TLockRequest
}

func GetNewEditorHub() *Hub {
	return &Hub{
		files:         make(map[string]*File),
		WriteRequest:  make(chan TWriteRequest),
		AddRequest:    make(chan TAddRequest),
		DeleteRequest: make(chan TDeleteRequest),
		LockRequest:   make(chan TLockRequest),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case request := <-hub.LockRequest:
			{
				file, ok := hub.files[request.FileName]
				if !ok {
					fmt.Println("Failed to find" + request.FileName)
					continue
				}
				old := file.UsedBy[request.Conn]
				file.UsedBy[request.Conn] = request.LockedLine
				for conn, _ := range file.UsedBy {
					if conn == request.Conn {
						continue
					}
					var content struct {
						Unlocked int    `json:"unlocked"`
						Locked   int    `json:"locked"`
						By       string `json:"by"`
					}
					content.Unlocked = old
					content.Locked = request.LockedLine
					content.By = request.By
					conn.WriteJSON(content)
				}
			}
		case request := <-hub.WriteRequest:
			{
				file, ok := hub.files[request.FileName]
				if !ok {
					fmt.Println("Failed to find" + request.FileName)
					continue
				}
				file.Content = request.Data
				for conn, _ := range file.UsedBy {
					if conn == request.Conn {
						continue
					}
					var content struct {
						Content string `json:"content"`
					}
					content.Content = request.Data
					conn.WriteJSON(content)
				}
			}
		case request := <-hub.AddRequest:
			{
				file, ok := hub.files[request.FileName]
				if !ok {
					content, err := readFileContent(request.FileName)
					var response TAddResponse
					if err != nil {
						fmt.Printf("Error reading file %v: %v\n", request.FileName, err)
						response.Error = err
						request.Response <- response
						continue
					}
					var newFile File
					newFile.FileName = request.FileName
					newFile.UsedBy = map[*websocket.Conn]int{
						request.Conn: 0,
					}
					newFile.Content = content
					hub.files[request.FileName] = &newFile
					response.Error = nil
					response.Content = content
					request.Response <- response
				} else {
					var response TAddResponse
					response.Content = file.Content
					response.Error = nil
					file.UsedBy[request.Conn] = 0
					request.Response <- response
				}
				fmt.Printf("Successfully added %p to file %v\n", request.Conn, request.FileName)
			}
		case request := <-hub.DeleteRequest:
			{
				file, ok := hub.files[request.FileName]
				if !ok {
					fmt.Println("Failed to find" + request.FileName)
					continue
				}
				fmt.Printf("Deleted %p from %v\n", request.Conn, request.FileName)
				delete(file.UsedBy, request.Conn)
			}
		}
	}
}
