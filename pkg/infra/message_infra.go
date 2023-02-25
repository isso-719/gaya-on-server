package infra

import (
	"context"
	model "github.com/isso-719/gaya-on-server/pkg/domain/model"
	"github.com/isso-719/gaya-on-server/pkg/domain/repository"
	"github.com/jinzhu/gorm"
)

type messageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) repository.IFMessageRepository {
	return &messageRepository{
		DB: db,
	}
}

func (s *messageRepository) Migrate() error {
	s.DB.AutoMigrate(&model.Message{})
	return nil
}

func (s *messageRepository) CreateMessage(ctx context.Context, roomID uint, messageType, messageBody string) error {
	tx := s.DB.BeginTx(ctx, nil)
	defer tx.RollbackUnlessCommitted()

	// Messageを作成
	message := model.Message{
		RoomID: roomID,
		Type:   messageType,
		Body:   messageBody,
	}
	if err := s.DB.Create(&message).Error; err != nil {
		return err
	}
	return nil
}

func (s *messageRepository) GetAllMessages(ctx context.Context, roomID uint) ([]*model.MessageTypeAndBody, error) {
	tx := s.DB.BeginTx(ctx, nil)
	defer tx.RollbackUnlessCommitted()

	// MessageTypeAndBodyを取得
	var messages []*model.MessageTypeAndBody
	if err := tx.Table("messages").Select("type, body").Where("room_id = ?", roomID).Scan(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
