package websocket

import "github.com/gorilla/websocket"

type Hub struct {
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
	set, ok := h.clients[c.UserID]
	if !ok{
		set = make(map[*Client]bool)
		h.clients[c.UserID] = set
	}
	set[c] = true
}

func (h *Hub) Unregister (c *Client) {
	
}