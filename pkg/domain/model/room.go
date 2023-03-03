package model

import (
	"golang.org/x/net/websocket"
	"time"
)

type Room struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	Token     string     `json:"token"` // Room Token
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// Websocket Client
type Client struct {
	RoomID int64
	Conn   *websocket.Conn
}

// Websocket Message
type WebSocketContent struct {
	RoomID int64
	Event  WebSocketEvent
}

const (
	WS_Connected    = "connected"
	WS_Disconnected = "disconnected"
	WS_Error        = "error"
	WS_Message      = "message"
)

// Websocket WebSocketEvent
type WebSocketEvent struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}
