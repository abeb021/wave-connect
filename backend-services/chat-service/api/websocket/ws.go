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
}

func NewWSHandler(srv *service.Service) *WSHandler {
	return &WSHandler{Srv: srv}
}

func (ws *WSHandler) ServeWS(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		http.Error(w, "upgrade failed", http.StatusBadRequest)
		return
	}
	defer conn.Close()
	userID := r.Header.Get("X-User-ID")
	//TODO: fix context parsing
	//userID := r.Context().Value("userID")
	log.Printf("User connected: %v", userID)

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil{
			break
		}

		err = conn.WriteMessage(mt, msg)
		if err != nil{
			break
		}
	}
}