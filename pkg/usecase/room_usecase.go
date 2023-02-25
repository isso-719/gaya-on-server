package usecase

import (
	"context"
	"crypto/rand"
	"errors"
	"github.com/isso-719/gaya-on-server/pkg/domain/service"
)

type IFRoomUsecase interface {
	Migrate() error
	CreateRoom(context.Context) (string, error)
	FindRoom(context.Context, string) (bool, error)
}

type roomUsecase struct {
	roomService service.IFRoomService
}

func NewRoomUsecase(ss service.IFRoomService) IFRoomUsecase {
	return &roomUsecase{
		roomService: ss,
	}
}

func (su *roomUsecase) Migrate() error {
	return su.roomService.Migrate()
}

func generateRandomToken(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func (su *roomUsecase) CreateRoom(ctx context.Context) (string, error) {
	var token string
	var err error
	// ランダムなトークンを生成して、roomIDとしてDBに保存する
	for {
		token, err = generateRandomToken(6)
		if err != nil {
			return "", err
		}
		ok, err := su.roomService.CreateRoom(ctx, token)
		if err != nil {
			return "", err
		}
		if ok {
			break
		}
	}
	return token, nil
}

func (su *roomUsecase) FindRoom(ctx context.Context, token string) (bool, error) {
	_, ok, err := su.roomService.FindRoom(ctx, token)
	if err != nil {
		return false, err
	}
	return ok, nil
}