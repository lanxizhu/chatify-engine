package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[string]*websocket.Conn)

var mu sync.Mutex

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
	Type    int    `json:"type"`
}

type WsHandler struct {
	clients map[string]*websocket.Conn
	mu      sync.Mutex
}

func SetupWsHandler() *WsHandler {
	return &WsHandler{
		clients: make(map[string]*websocket.Conn),
	}
}

type Client struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (ws *WsHandler) Connect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	var client Client
	if err := c.ShouldBindUri(&client); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	clientIP := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	log.Printf("Client IP: %s, User-Agent: %s", clientIP, userAgent)
	clientInfo := clientIP + " " + userAgent
	println(clientInfo)

	mu.Lock()
	clients[client.ID] = conn
	mu.Unlock()

	conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Connection closed: %s, Code: %d, Text: %s, now clients count: %v", client.ID, code, text, len(clients)-1)
		return nil
	})

	conn.SetPingHandler(func(appData string) error {
		log.Printf("Received ping: %s", appData)
		s := "pong"
		if err = conn.WriteMessage(websocket.PongMessage, []byte(s)); err != nil {
			log.Println("Error writing pong message:", err)
			return err
		}
		return nil
	})

	defer func() {
		mu.Lock()
		delete(clients, client.ID)
		mu.Unlock()
		log.Printf("Client disconnected: %s, now clients count: %v", client.ID, len(clients))
	}()

	for {
		messageType, p, readErr := conn.ReadMessage()
		if readErr != nil {
			log.Println(readErr)
			return
		}

		if string(p) == "ping" {
			s := "pong"
			if err = conn.WriteMessage(websocket.TextMessage, []byte(s)); err != nil {
				log.Println("Error writing pong message:", err)
			}
			continue
		}

		println("messageType: ", messageType, string(p) == "ping")

		var message Message
		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Println(err)
			return
		}

		var targetConn *websocket.Conn
		var content string

		if message.Type == websocket.PingMessage {
			targetConn = conn
			content = "pong"
			//messageType = websocket.PongMessage
		} else if message.Type == websocket.TextMessage {
			targetConn = clients[message.To]
			content = message.Content

			if targetConn == nil {
				targetConn = conn
				content = "User not found"
				//messageType = websocket.PingMessage
			}
		} else {
			targetConn = conn
			content = "Invalid message type"
		}

		println("content: ", content)

		if err = targetConn.WriteMessage(messageType, []byte(content)); err != nil {
			//log.Println(err)
			return
		}
	}
}

func SendToClient(id string, connect string) {
	conn := clients[id]
	if conn != nil {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(connect)); err != nil {
			log.Println(err)
			return
		}
	}
}
