package repository

import (
	"context"
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
)

type IFRoomRepository interface {
	Migrate() error
	CreateRoom(ctx context.Context, token string) (bool, error)
	FindRoom(ctx context.Context, token string) (*model.Room, bool, error)
}
