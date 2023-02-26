package service

import (
	"context"
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
	"github.com/isso-719/gaya-on-server/pkg/domain/repository"
)

type IFMessageService interface {
	Migrate() error
	CreateMessage(ctx context.Context, roomID int64, messageType, messageBody string) error
	GetAllMessages(ctx context.Context, roomID int64) ([]*model.MessageTypeAndBody, error)
}

type messageService struct {
	messageRepository repository.IFMessageRepository
}

func NewMessageService(
	sr repository.IFMessageRepository) IFMessageService {
	return &messageService{
		messageRepository: sr,
	}
}

func (ss *messageService) Migrate() error {
	return ss.messageRepository.Migrate()
}

func (ss *messageService) CreateMessage(ctx context.Context, roomID int64, messageType, messageBody string) error {
	return ss.messageRepository.CreateMessage(ctx, roomID, messageType, messageBody)
}

func (ss *messageService) GetAllMessages(ctx context.Context, roomID int64) ([]*model.MessageTypeAndBody, error) {
	return ss.messageRepository.GetAllMessages(ctx, roomID)
}
