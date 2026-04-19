package websocket

import (
	"chat-service/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSHandler struct {
	Srv *service.Service
	Hub *Hub
}

func NewWSHandler(srv *service.Service, hub *Hub) *WSHandler {
	return &WSHandler{Srv: srv, Hub: hub}
}

func (ws *WSHandler) ServeWS(w http.ResponseWriter, r *http.Request){
	userID := r.Header.Get("X-User-ID")
	if userID == ""{
		http.Error(w, "missing user id", http.StatusUnauthorized)
		return
	}

	// upgrade http to ws
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		http.Error(w, "upgrade failed", http.StatusBadRequest)
		return
	}

	// create client of a user (phone, laptop, browser)
	c := &Client{
		Conn:   conn,
		UserID: userID,
		Send:   make(chan []byte),
		Hub:    ws.Hub,
	}

	//register this single connection to a hub (map[UserId]map[*Client]bool)
	ws.Hub.Register(c)
	log.Printf("User connected: %v", userID)

	//initial message
	c.Send <- []byte("connected")

	//start both
	go c.writePump()
	//block till return
	c.readPump()
}
