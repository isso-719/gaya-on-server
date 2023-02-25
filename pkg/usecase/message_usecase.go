package usecase

import (
	"context"
	"errors"
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
	"github.com/isso-719/gaya-on-server/pkg/domain/service"
)

type IFMessageUsecase interface {
	Migrate() error
	CreateMessage(ctx context.Context, token, messageType, messageBody string) error
	GetAllMessages(ctx context.Context, token string) ([]*model.MessageTypeAndBody, error)
}

type messageUsecase struct {
	messageService service.IFMessageService
	roomService    service.IFRoomService
}

func NewMessageUsecase(
	ss service.IFMessageService,
	roomService service.IFRoomService,
) IFMessageUsecase {
	return &messageUsecase{
		messageService: ss,
		roomService:    roomService,
	}
}

func (su *messageUsecase) Migrate() error {
	return su.messageService.Migrate()
}

func (su *messageUsecase) CreateMessage(ctx context.Context, token, messageType, messageBody string) error {
	if messageType != model.MessageTypeText && messageType != model.MessageTypeEmoji {
		return errors.New("messageType is invalid")
	}
	if messageBody == "" {
		return errors.New("messageBody is empty")
	}
	room, ok, err := su.roomService.FindRoom(ctx, token)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("room not found")
	}
	return su.messageService.CreateMessage(ctx, room.ID, messageType, messageBody)
}

func (su *messageUsecase) GetAllMessages(ctx context.Context, token string) ([]*model.MessageTypeAndBody, error) {
	room, ok, err := su.roomService.FindRoom(ctx, token)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("room not found")
	}
	return su.messageService.GetAllMessages(ctx, room.ID)
}
