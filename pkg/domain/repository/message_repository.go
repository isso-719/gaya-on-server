package repository

import (
	"context"
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
)

type IFMessageRepository interface {
	Migrate() error
	CreateMessage(ctx context.Context, roomID uint, messageType, messageBody string) error
	GetAllMessages(ctx context.Context, roomID uint) ([]*model.MessageTypeAndBody, error)
}
