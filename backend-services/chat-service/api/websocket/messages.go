package websocket

import (
	"encoding/json"
	"time"
)

const (
	TypeChatSend    = "chat.send"
	TypeChatMessage = "chat.message"
	TypeChatError   = "chat.error"
)

type inboundMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type chatSendPayload struct {
	Receiver string `json:"receiver"`
	Text     string `json:"text"`
}

type outboundMessage struct {
	Type      string    `json:"type"`
	TimeStamp time.Time `json:"data,omitempty"`
	ID        string    `json:"id,omitempty"`
	Payload   any       `json:"payload,omitempty"`
	Err       string    `json:"error,omitempty"`
}
