package websocket

import (
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMsgSize = 64 * 1024
)


func (c *Client) readPump () {
	defer func(){
		c.Hub.Unregister(c)
		close(c.Send)
		_ = c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMsgSize)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error{
		return c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, msg, err := c.Conn.ReadMessage()

		if err != nil{
			var closeErr *websocket.CloseError
			if errors.As(err, &closeErr){
				return
			}
			return
		}

		//for now just echo back
		select {
		case c.Send <- msg:
		default:
			return
		}
	}
}

func (c *Client) writePump () {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.Conn.Close()
	}()



	for {
		select{
		case msg, ok := <-c.Send:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			} 
			if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil{
				return
			}

		case <- ticker.C:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil{
				log.Printf("ws ping error user=%s: %v", c.UserID, err)
				return
			}
		}
	}
}