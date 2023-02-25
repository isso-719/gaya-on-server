package infra

import (
	"context"
	model "github.com/isso-719/gaya-on-server/pkg/domain/model"
	"github.com/isso-719/gaya-on-server/pkg/domain/repository"
	"github.com/jinzhu/gorm"
)

type roomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) repository.IFRoomRepository {
	return &roomRepository{
		DB: db,
	}
}

func (s *roomRepository) Migrate() error {
	s.DB.AutoMigrate(&model.Room{})
	return nil
}

func (s *roomRepository) CreateRoom(ctx context.Context, token string) (bool, error) {
	tx := s.DB.BeginTx(ctx, nil)
	defer tx.RollbackUnlessCommitted()

	var room model.Room
	if err := tx.Where("token = ?", token).First(&room).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 重複していない場合は作成
			room := model.Room{
				Token: token,
			}
			if err := tx.Create(&room).Error; err != nil {
				return false, err
			}
			tx.Commit()
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (s *roomRepository) FindRoom(ctx context.Context, token string) (*model.Room, bool, error) {
	tx := s.DB.BeginTx(ctx, nil)
	defer tx.RollbackUnlessCommitted()

	var room model.Room
	if err := tx.Where("token = ?", token).First(&room).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &room, true, nil
}
