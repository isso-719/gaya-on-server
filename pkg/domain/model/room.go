package model

import (
	"golang.org/x/net/websocket"
	"time"
)

type Room struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Token     string     `json:"token"` // Room Token
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// Websocket Client
type Client struct {
	RoomID uint
	Conn   *websocket.Conn
}

// Websocket Message
type WebSocketContent struct {
	RoomID uint
	Event  WebSocketEvent
}

// Websocket WebSocketEvent
type WebSocketEvent struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}
