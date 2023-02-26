package repository

import (
	"context"
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
)

type IFMessageRepository interface {
	Migrate() error
	CreateMessage(ctx context.Context, roomID int64, messageType, messageBody string) error
	GetAllMessages(ctx context.Context, roomID int64) ([]*model.MessageTypeAndBody, error)
}
