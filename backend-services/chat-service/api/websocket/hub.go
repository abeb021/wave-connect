package websocket

import (
	"sync"
)

type Hub struct {
	mu sync.RWMutex
	//list of connections of a user (userID)
	clients map[string]map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Register(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	set, ok := h.clients[c.UserID]
	//if user is not registred then we need to create him a place by his id
	if !ok {
		set = make(map[*Client]bool)
		h.clients[c.UserID] = set
	}
	// so if its his 2nd or more connection we just add it
	set[c] = true
}

func (h *Hub) Unregister(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	set, ok := h.clients[c.UserID]
	if !ok {
		return
	}
	//we delete only this connection
	delete(set, c)
	//if it was the last connection we delete user from clients map
	if len(set) == 0 {
		delete(h.clients, c.UserID)
	}
}

func (h *Hub) SendToUser(userID string, msg []byte) {
	h.mu.RLock()
	set, ok := h.clients[userID]
	if !ok {
		h.mu.RUnlock()
		return
	}

	clients := make([]*Client, 0, len(set))
	for c := range set {
		clients = append(clients, c)
	}
	h.mu.RUnlock()

	for _, c := range clients {
		select {
		case c.Send <- msg:
		default:
			h.Unregister(c)
			_ = c.Conn.Close()
		}
	}
}
