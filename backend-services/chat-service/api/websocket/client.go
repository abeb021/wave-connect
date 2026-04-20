package websocket

import (
	"chat-service/internal/repository"
	"chat-service/internal/service"
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	UserID string
	Send   chan []byte
	Hub    *Hub
	Srv    *service.Service
}

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMsgSize = 64 * 1024
)

func (c *Client) readPump() {
	defer func() {
		c.Hub.Unregister(c)
		close(c.Send)
	}()

	c.Conn.SetReadLimit(maxMsgSize)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		return c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, msg, err := c.Conn.ReadMessage()

		if err != nil {
			var closeErr *websocket.CloseError
			if errors.As(err, &closeErr) {
				return
			}
			return
		}

		go c.handleMessage(msg)
	}
}

func (c *Client) handleMessage(msg []byte) {
	var in inboundMessage
	if err := json.Unmarshal(msg, &in); err != nil {
		c.sendError("invalid json", "")
		return
	}

	switch in.Type {
	case TypeChatSend:
		var p chatSendPayload
		if err := json.Unmarshal(in.Payload, &p); err != nil {
			c.sendError("invalid chat payload", "")
			return
		}

		if p.Receiver == "" || p.Text == "" {
			c.sendError("receiver and text required", "")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		saved, err := c.Srv.CreateMessage(ctx, &repository.MessageRequest{
			Text:     p.Text,
			Receiver: p.Receiver,
			Sender:   c.UserID,
		})
		if err != nil {
			c.sendError("failed to save message", "")
			return
		}

		resp := outboundMessage{
			Type:      TypeChatMessage,
			TimeStamp: saved.TimeSent,
			ID:        saved.ID,
			Payload:   p,
		}

		out, err := json.Marshal(resp)
		if err != nil {
			return
		}

		c.Hub.SendToUser(p.Receiver, out)
		c.Hub.SendToUser(c.UserID, out)

	default:
		c.sendError("unknown message type", "")
	}
}

func (c *Client) sendError(msg, id string) {
	c.queueJson(outboundMessage{
		Type:      TypeChatError,
		TimeStamp: time.Now(),
		ID:        id,
		Err:       msg,
	})
}

func (c *Client) queueJson(out outboundMessage) {
	b, err := json.Marshal(out)
	if err != nil {
		return
	}

	select {
	case c.Send <- b:
	default:
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			//if for some reason the channel is closed, we close the ws connection
			if !ok {
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//if ok we send received message from the channel
			if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}

		// for ping pong ws
		case <-ticker.C:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("ws ping error user=%s: %v", c.UserID, err)
				return
			}
		}
	}
}
