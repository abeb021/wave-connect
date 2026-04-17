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
	set, ok := h.clients[c.UserID]
	if !ok {
		return
	}
	delete(set, c)
	if len(set) == 0{
		delete(h.clients, c.UserID)
	}
}

func (h *Hub) SendToUser (c *Client, msg []byte){
	set := h.clients[c.UserID]

	for c := range set{
		select {
		case c.Send <- msg:
		default:
			h.Unregister(c)
			_ = c.Conn.Close()
		}
	}
}

//TODO add mutexes