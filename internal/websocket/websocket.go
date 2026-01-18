package websocket

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"main/internal/utils"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ClientManager maintains the set of active clients and broadcasts messages to the clients.
type ClientManager struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mutex      sync.RWMutex
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	manager *ClientManager
	socket  *websocket.Conn
	send    chan []byte
}

// Manager is the global instance of ClientManager
var Manager = ClientManager{
	clients:    make(map[*Client]bool),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	broadcast:  make(chan []byte),
}

func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-manager.register:
			manager.mutex.Lock()
			manager.clients[client] = true
			manager.mutex.Unlock()
			log.Println("New WebSocket client connected")
		case client := <-manager.unregister:
			manager.mutex.Lock()
			if _, ok := manager.clients[client]; ok {
				close(client.send)
				delete(manager.clients, client)
			}
			manager.mutex.Unlock()
			log.Println("WebSocket client disconnected")
		case message := <-manager.broadcast:
			manager.mutex.RLock()
			for client := range manager.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(manager.clients, client)
				}
			}
			manager.mutex.RUnlock()
		}
	}
}

func (manager *ClientManager) BroadcastNotification(notification models.Notification) {
	data1 := map[string]interface{}{
		"id":           notification.ID,
		"title":        notification.Title,
		"message":      notification.Message,
		"type":         notification.Type,
		"reference_id": notification.ReferenceID,
		"created_at":   notification.CreatedAt,
		"is_read":      notification.IsRead,
	}

	data, err := json.Marshal(data1)
	if err != nil {
		log.Println("Error marshalling notification:", err)
		return
	}
	manager.broadcast <- data
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func (c *Client) readPump() {
	defer func() {
		c.manager.unregister <- c
		c.socket.Close()
	}()
	for {
		_, _, err := c.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (c *Client) writePump() {
	defer func() {
		c.socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// ServeWS handles websocket requests from the peer.
func ServeWS(c *gin.Context) {
	// Validate Token from Query Param
	token := c.Query("token")
	log.Printf("WebSocket connection attempt. Token received: %s", token) // Debug log

	if token == "" {
		log.Println("WebSocket connection failed: Token missing")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
		return
	}

	userID, role, err := utils.ValidateJWT(token)
	if err != nil {
		log.Printf("WebSocket connection failed: Invalid token. Error: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid token",
			"details": err.Error(),
		})
		return
	}

	log.Printf("WebSocket authorized for UserID: %d, Role: %s", userID, role)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{manager: &Manager, socket: conn, send: make(chan []byte, 256)}
	client.manager.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
