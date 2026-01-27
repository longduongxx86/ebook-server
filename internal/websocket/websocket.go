package websocket

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"main/internal/repositories"
	"main/internal/utils"
	"net/http"
	"sync"
	"time"

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
	UserID  uint
	Role    string
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
			log.Printf("New WebSocket client connected. UserID: %d, Role: %s", client.UserID, client.Role)
		case client := <-manager.unregister:
			manager.mutex.Lock()
			if _, ok := manager.clients[client]; ok {
				close(client.send)
				delete(manager.clients, client)
			}
			manager.mutex.Unlock()
			log.Printf("WebSocket client disconnected. UserID: %d", client.UserID)
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

func (manager *ClientManager) SendToUser(userID uint, message []byte) {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()
	for client := range manager.clients {
		if client.UserID == userID {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(manager.clients, client)
			}
		}
	}
}

func (manager *ClientManager) SendToAdmins(message []byte) {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()
	for client := range manager.clients {
		if client.Role == "admin" || client.Role == "manager" {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(manager.clients, client)
			}
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

	data, err := json.Marshal(map[string]interface{}{
		"type":    "notification",
		"payload": data1,
	})
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

type WSMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type ChatPayload struct {
	Content  string `json:"content"`
	ToUserID uint   `json:"to_user_id,omitempty"`
}

func (c *Client) readPump() {
	defer func() {
		c.manager.unregister <- c
		c.socket.Close()
	}()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var msg WSMessage
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Invalid JSON:", err)
			continue
		}

		if msg.Type == "chat" {
			var payload ChatPayload
			if err := json.Unmarshal(msg.Payload, &payload); err != nil {
				log.Println("Invalid Chat Payload:", err)
				continue
			}
			c.handleChatMessage(payload)
		}
	}
}

func (c *Client) handleChatMessage(payload ChatPayload) {
	chatRepo := repositories.NewChatRepository()

	var conversation *models.Conversation
	var err error
	var conversationID uint

	// Determine Conversation
	if c.Role == "customer" || c.Role == "" {
		// Customer sending message: conversation is their own
		conversation, err = chatRepo.FirstOrCreateConversation(c.UserID)
		if err != nil {
			log.Println("Error getting conversation:", err)
			return
		}
		conversationID = conversation.ID
	} else {
		// Admin sending message: must specify recipient
		if payload.ToUserID == 0 {
			return
		}
		conversation, err = chatRepo.FirstOrCreateConversation(payload.ToUserID)
		if err != nil {
			log.Println("Error getting conversation:", err)
			return
		}
		conversationID = conversation.ID
	}

	// Save Message
	msg := models.Message{
		ConversationID: conversationID,
		SenderID:       c.UserID,
		Content:        payload.Content,
		CreatedAt:      time.Now().Unix(),
		IsAdmin:        c.Role == "admin" || c.Role == "manager",
	}

	if err := chatRepo.CreateMessage(&msg); err != nil {
		log.Println("Error saving message:", err)
		return
	}

	// Update Conversation Last Message
	if err := chatRepo.UpdateConversationLastMessage(conversation.ID, payload.Content, msg.CreatedAt); err != nil {
		log.Println("Error updating conversation:", err)
	}

	// Load Sender Info for frontend display
	fullMsg, err := chatRepo.GetMessageWithSender(msg.ID)
	if err != nil {
		log.Println("Error loading message details:", err)
		return
	}

	// Prepare Response
	response := map[string]interface{}{
		"type":    "chat",
		"payload": fullMsg,
	}
	respBytes, _ := json.Marshal(response)

	// Route Message
	// Send to the conversation owner (the customer)
	Manager.SendToUser(conversation.UserID, respBytes)

	// Send to all admins
	Manager.SendToAdmins(respBytes)
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
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
		return
	}

	userID, role, err := utils.ValidateJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		manager: &Manager,
		socket:  conn,
		send:    make(chan []byte, 256),
		UserID:  userID,
		Role:    role,
	}
	client.manager.register <- client

	go client.writePump()
	go client.readPump()
}
