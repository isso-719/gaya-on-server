package model

import "time"

const (
	MessageTypeText  = "text"
	MessageTypeEmoji = "emoji"
)

type Message struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	RoomID    int64      `json:"room_id"` // Room ID
	Type      string     `json:"type"`    // Message Type (text, emoji)
	Body      string     `json:"body"`    // Message Body
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type MessageTypeAndBody struct {
	Type string `json:"type"`
	Body string `json:"body"`
}
