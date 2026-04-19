package websocket

import (
	"sync"
	"github.com/gorilla/websocket"
)

type Hub struct {
	mu      sync.RWMutex
	clients map[string]map[*Client]bool
}

type Client struct {
    Conn   *websocket.Conn
    UserID string
    Send   chan []byte
    Hub    *Hub
}

func NewHub() *Hub{
	return &Hub{
		clients: make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Register (c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	set, ok := h.clients[c.UserID]
	if !ok{
		set = make(map[*Client]bool)
		h.clients[c.UserID] = set
	}
	set[c] = true
}

func (h *Hub) Unregister (c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	set, ok := h.clients[c.UserID]
	if !ok {
		return
	}
	delete(set, c)
	if len(set) == 0{
		delete(h.clients, c.UserID)
	}
}

func (h *Hub) SendToUser (userID string, msg []byte){
	h.mu.RLock()
	set := h.clients[userID]
	defer h.mu.RUnlock()

	for c := range set{
		select {
		case c.Send <- msg:
		default:
			h.Unregister(c)
			_ = c.Conn.Close()
		}
	}
}