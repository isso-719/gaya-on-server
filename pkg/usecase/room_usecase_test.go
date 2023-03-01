package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	test_helper "github.com/isso-719/gaya-on-server/lib/test"
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
	"github.com/isso-719/gaya-on-server/pkg/domain/service"
	"strings"
	"testing"
	"time"
)

// TestGenerateRandomToken : generateRandomTokenのテスト、指定された文字列の長さと規定の文字列のみを含むかを確認する
func TestGenerateRandomToken(t *testing.T) {
	//token, err := generateRandomToken(6)
	//if err != nil {
	//	t.Error(err)
	//}
	//if len(token) != 6 {
	//	t.Error("token length is invalid")
	//}
	//
	//const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	//for _, v := range token {
	//	if !strings.Contains(letters, string(v)) {
	//		t.Error("token contains invalid character")
	//	}
	//}

	type Args struct {
		digit uint32
	}
	type Returns struct {
		token *string
		err   error
	}
	type testContext struct {
		args    Args
		returns Returns
	}

	tests := []struct {
		name        string
		testContext func(ctrl *gomock.Controller) *testContext
	}{
		{
			name: "正常, トークンを生成して返す",
			testContext: func(ctrl *gomock.Controller) *testContext {
				// 6 文字のトークン
				token := "123456"
				return &testContext{
					args: Args{
						digit: 6,
					},
					returns: Returns{
						token: &token,
						err:   nil,
					},
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := tt.testContext(ctrl)
			token, err := generateRandomToken(ctx.args.digit)
			if err != nil {
				t.Error(err)
			}
			// 使用できない文字が含まれていないか確認
			const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
			for _, v := range token {
				if !strings.Contains(letters, string(v)) {
					t.Error("token contains invalid character")
				}
			}
			// トークンの長さが正しいか確認
			if len(token) != int(ctx.args.digit) {
				t.Error("token length is invalid")
			}
		})
	}
}

func TestCreateRoom(t *testing.T) {
	type Fields struct {
		roomService service.IFRoomService
	}
	type Args struct {
		ctx context.Context
	}
	type Returns struct {
		token *string
		err   error
	}
	type testContext struct {
		fields  Fields
		args    Args
		returns Returns
	}

	tests := []struct {
		name        string
		testContext func(ctrl *gomock.Controller) *testContext
	}{
		{
			name: "正常, 1 回目はエラーが発生し、2 回目は成功する",
			testContext: func(ctrl *gomock.Controller) *testContext {
				roomService := service.NewMockIFRoomService(ctrl)
				roomService.EXPECT().CreateRoom(gomock.Any(), gomock.Any()).Return(false, errors.New("error")).Times(1)
				roomService.EXPECT().CreateRoom(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)

				token := "1a2b3c"
				return &testContext{
					fields: Fields{
						roomService: roomService,
					},
					args: Args{
						ctx: context.Background(),
					},
					returns: Returns{
						token: &token,
						err:   nil,
					},
				}
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			test_helper.RunTestWithGoMock(t, func(ctrl *gomock.Controller) {
				tc := tt.testContext(ctrl)
				u := &roomUsecase{
					roomService: tc.fields.roomService,
				}

				// 1 回目はエラー
				token, err := u.CreateRoom(tc.args.ctx)
				if err == nil {
					t.Error("error should be returned")
				}
				if token != nil {
					t.Error("token should be nil")
				}

				// 2 回目は正常にトークンを生成する
				token, err = u.CreateRoom(tc.args.ctx)
				if err != nil {
					t.Error(err)
				}
				if token == nil {
					t.Error("token should not be nil")
				}

				// トークンの長さが正しいか確認
				if len(*token) != 6 {
					t.Error("token length is invalid")
				}
			})
		})
	}
}

func TestFindRoom(t *testing.T) {
	type Fields struct {
		roomService service.IFRoomService
	}
	type Args struct {
		ctx    context.Context
		roomId string
	}
	type Returns struct {
		ok  bool
		err error
	}
	type testContext struct {
		fields  Fields
		args    Args
		returns Returns
	}

	tests := []struct {
		name        string
		testContext func(ctrl *gomock.Controller) *testContext
	}{
		{
			name: "正常, ルームが存在する",
			testContext: func(ctrl *gomock.Controller) *testContext {
				roomService := service.NewMockIFRoomService(ctrl)
				roomService.EXPECT().FindRoom(gomock.Any(), gomock.Any()).Return(
					&model.Room{
						ID:        1,
						Token:     "a1b2c3",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
						DeletedAt: nil,
					},
					true,
					nil,
				).Times(1)

				return &testContext{
					fields: Fields{
						roomService: roomService,
					},
					args: Args{
						ctx:    context.Background(),
						roomId: "1a2b3c",
					},
					returns: Returns{
						ok:  true,
						err: nil,
					},
				}
			},
		},
		{
			name: "異常, ルームが存在しない",
			testContext: func(ctrl *gomock.Controller) *testContext {
				roomService := service.NewMockIFRoomService(ctrl)
				roomService.EXPECT().FindRoom(gomock.Any(), gomock.Any()).Return(
					nil,
					false,
					errors.New("room not found"),
				).Times(1)

				return &testContext{
					fields: Fields{
						roomService: roomService,
					},
					args: Args{
						ctx:    context.Background(),
						roomId: "1a2b3c",
					},
					returns: Returns{
						ok:  false,
						err: errors.New("room not found"),
					},
				}
			},
		},
		{
			name: "異常, token が空文字",
			testContext: func(ctrl *gomock.Controller) *testContext {
				roomService := service.NewMockIFRoomService(ctrl)
				roomService.EXPECT().FindRoom(gomock.Any(), gomock.Any()).Return(
					nil,
					false,
					errors.New("room not found"),
				).Times(1)

				return &testContext{
					fields: Fields{
						roomService: roomService,
					},
					args: Args{
						ctx:    context.Background(),
						roomId: "1a2b3c",
					},
					returns: Returns{
						ok:  false,
						err: errors.New("room not found"),
					},
				}
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			test_helper.RunTestWithGoMock(t, func(ctrl *gomock.Controller) {
				tc := tt.testContext(ctrl)
				s := &roomUsecase{
					roomService: tc.fields.roomService,
				}
				ok, err := s.FindRoom(tc.args.ctx, tc.args.roomId)
				if err != nil && err.Error() != tc.returns.err.Error() {
					t.Error("expected err is", tc.returns.err, "but actual is", err)
				}
				if ok != tc.returns.ok {
					t.Error("expected ok is", tc.returns.ok, "but actual is", ok)
				}
			})
		})
	}
}

// TODO: implement me
func TestJoinRoom(t *testing.T) {

}
