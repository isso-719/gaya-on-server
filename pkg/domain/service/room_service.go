package service

import (
	"context"
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
	"github.com/isso-719/gaya-on-server/pkg/domain/repository"
)

type IFRoomService interface {
	Migrate() error
	CreateRoom(ctx context.Context, token string) (bool, error)
	FindRoom(ctx context.Context, token string) (*model.Room, bool, error)
}

type roomService struct {
	roomRepository repository.IFRoomRepository
}

func NewRoomService(sr repository.IFRoomRepository) IFRoomService {
	return &roomService{
		roomRepository: sr,
	}
}
func (ss *roomService) Migrate() error {
	return ss.roomRepository.Migrate()
}

func (ss *roomService) CreateRoom(ctx context.Context, token string) (bool, error) {
	return ss.roomRepository.CreateRoom(ctx, token)
}

func (ss *roomService) FindRoom(ctx context.Context, token string) (*model.Room, bool, error) {
	return ss.roomRepository.FindRoom(ctx, token)
}
