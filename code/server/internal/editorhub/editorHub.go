package editorhub

import (
	"fmt"
	"slices"

	"github.com/gorilla/websocket"
)

type File struct {
	FileName string
	Content  [][]byte
	UsedBy   map[*websocket.Conn]bool
}

type TDifference struct {
	Index int    `json:"index"`
	Data  string `json:"data"`
}

type TWriteRequest struct {
	FileName   string
	Conn       *websocket.Conn
	Difference []TDifference
}

type TAddResponse struct {
	Error   error
	Content [][]byte
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
}

func GetNewEditorHub() *Hub {
	return &Hub{
		files:         make(map[string]*File),
		WriteRequest:  make(chan TWriteRequest),
		AddRequest:    make(chan TAddRequest),
		DeleteRequest: make(chan TDeleteRequest),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case request := <-hub.WriteRequest:
			{
				file, ok := hub.files[request.FileName]
				if !ok {
					fmt.Println("Failed to find" + request.FileName)
					continue
				}
				slices.SortFunc(request.Difference, func(prev TDifference, next TDifference) int {
					return prev.Index - next.Index
				})
				for _, diff := range request.Difference {
					if diff.Index >= len(file.Content) {
						file.Content = append(file.Content, []byte(diff.Data))
					} else {
						file.Content[diff.Index] = []byte(diff.Data)
					}
				}
				for conn, _ := range file.UsedBy {
					if conn == request.Conn {
						continue
					}
					var diff struct {
						Difference []TDifference `json:"diff"`
					}
					diff.Difference = request.Difference
					conn.WriteJSON(diff)
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
					newFile.UsedBy = map[*websocket.Conn]bool{
						request.Conn: true,
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
					file.UsedBy[request.Conn] = true
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
