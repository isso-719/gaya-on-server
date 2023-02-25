package model

import (
	"time"
)

type Room struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Token     string     `json:"token"` // Room Token
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
