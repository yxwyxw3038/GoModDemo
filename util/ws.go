package util

import (
	"GoModDemo/model"
	"encoding/json"
	"github.com/gorilla/websocket"
)

// ClientManager is a websocket manager
type WSClientManager struct {
	Clients    map[*WSClient]bool
	Broadcast  chan []byte
	Register   chan *WSClient
	Unregister chan *WSClient
}

// Client is a websocket client
type WSClient struct {
	ID     string
	UserId string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is an object for websocket message which is mapped to json type
type WSMessage struct {
	// Sender    string `json:"sender,omitempty"`
	// Recipient string `json:"recipient,omitempty"`
	// Content   string `json:"content,omitempty"`
	Data string `json:"Data,omitempty"`
	Type string `json:"Type,omitempty"`
	ID   string `json:"ID,omitempty"`
}

// Manager define a ws server manager
var WSManager = WSClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *WSClient),
	Unregister: make(chan *WSClient),
	Clients:    make(map[*WSClient]bool),
}

// Start is to start a ws server
func (manager *WSClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&WSMessage{ID: conn.ID, Data: "注册WS链接成功."})
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&WSMessage{ID: conn.ID, Data: "取消WS链接成功."})
				manager.Send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

// Send is to send ws message to ws client
func (manager *WSClientManager) Send(message []byte, ignore *WSClient) {
	for conn := range manager.Clients {
		if conn == ignore {
			conn.Send <- message
		}
	}
}

func (c *WSClient) Read() {
	defer func() {
		WSManager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			WSManager.Unregister <- c
			c.Socket.Close()
			break
		}
		var wsInfo model.WsInfoModel
		err = json.Unmarshal([]byte(message), &wsInfo)
		if err != nil {
			WSManager.Unregister <- c
			c.Socket.Close()
			break
		}
		dataMsg := ""
		Type := "99"
		switch wsInfo.Type {
		case "Register":
			dataMsg = wsInfo.ID
			c.UserId = wsInfo.ID

			break
		default:
			dataMsg = ""

		}
		Type = wsInfo.Type
		jsonMessage, _ := json.Marshal(&WSMessage{ID: c.ID, Type: Type, Data: dataMsg})
		WSManager.Broadcast <- jsonMessage
	}
}

func (c *WSClient) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
