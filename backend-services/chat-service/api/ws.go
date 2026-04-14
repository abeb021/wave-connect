package api

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		http.Error(w, "upgrade failed", http.StatusBadRequest)
	}
	defer conn.Close()

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