package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	test_helper "github.com/isso-719/gaya-on-server/lib/test"
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
	"github.com/isso-719/gaya-on-server/pkg/domain/service"
	"reflect"
	"testing"
	"time"
)

func TestCreateMessage(t *testing.T) {
	type Fields struct {
		roomService    service.IFRoomService
		messageService service.IFMessageService
	}
	type Args struct {
		ctx         context.Context
		token       string
		messageType string
		messageBody string
	}
	type Returns struct {
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
			name: "正常, text メッセージを作成できる",
			testContext: func(ctrl *gomock.Controller) *testContext {
				roomService := service.NewMockIFRoomService(ctrl)
				roomService.EXPECT().FindRoom(gomock.Any(), gomock.Any()).Return(
					&model.Room{
						ID:        1,
						Token:     "a1b2c3",
						CreatedAt: time.Now(),
					},
					true,
					nil,
				).Times(1)

				messageService := service.NewMockIFMessageService(ctrl)
				messageService.EXPECT().CreateMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

				return &testContext{
					fields: Fields{
						roomService:    roomService,
						messageService: messageService,
					},
					args: Args{
						ctx:         context.Background(),
						token:       "a1b2c3",
						messageType: "text",
						messageBody: "Hello, world!",
					},
					returns: Returns{
						err: nil,
					},
				}
			},
		},
		// TODO: 絵文字メッセージの追加テストの実装, 絵文字が一文字でないときは受け入れない時の実装=> 1 文字はなし
		//{
		//	name: "正常, emoji メッセージを作成できる",
		//	testContext: func(ctrl *gomock.Controller) *testContext {
		//		roomService := service.NewMockIFRoomService(ctrl)
		//		roomService.EXPECT().FindRoom(gomock.Any(), gomock.Any()).Return(
		//			&model.Room{
		//				ID:        1,
		//				Token:     "a1b2c3",
		//				CreatedAt: time.Now(),
		//			},
		//			true,
		//			nil,
		//		).Times(1)
		//
		//		messageService := service.NewMockIFMessageService(ctrl)
		//		messageService.EXPECT().CreateMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
		//
		//		return &testContext{
		//			fields: Fields{
		//				roomService:    roomService,
		//				messageService: messageService,
		//			},
		//			args: Args{
		//				ctx:         context.Background(),
		//				token:       "a1b2c3",
		//				messageType: "emoji",
		//				messageBody: "👍",
		//			},
		//			returns: Returns{
		//				err: nil,
		//			},
		//		}
		//	},
		//},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			test_helper.RunTestWithGoMock(t, func(ctrl *gomock.Controller) {
				tc := tt.testContext(ctrl)
				u := &messageUsecase{
					roomService:    tc.fields.roomService,
					messageService: tc.fields.messageService,
				}
				err := u.CreateMessage(tc.args.ctx, tc.args.token, tc.args.messageType, tc.args.messageBody)
				if !reflect.DeepEqual(err, tc.returns.err) {
					t.Errorf("CreateMessage() error = %v, wantErr %v", err, tc.returns.err)
				}
			})
		})
	}
}

func TestGetAllMessages(t *testing.T) {
	type Fields struct {
		roomService    service.IFRoomService
		messageService service.IFMessageService
	}
	type Args struct {
		ctx   context.Context
		token string
	}
	type Returns struct {
		messages []*model.MessageTypeAndBody
		err      error
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
			name: "正常, メッセージを取得できる",
			testContext: func(ctrl *gomock.Controller) *testContext {
				roomService := service.NewMockIFRoomService(ctrl)
				roomService.EXPECT().FindRoom(gomock.Any(), gomock.Any()).Return(
					&model.Room{
						ID:        1,
						Token:     "a1b2c3",
						CreatedAt: time.Now(),
					},
					true,
					nil,
				).Times(1)

				messageService := service.NewMockIFMessageService(ctrl)
				messages := []*model.MessageTypeAndBody{
					{
						Type: "text",
						Body: "Hello, world!",
					},
					{
						Type: "emoji",
						Body: "🥳",
					},
				}
				messageService.EXPECT().GetAllMessages(gomock.Any(), gomock.Any()).Return(
					messages,
					nil,
				).Times(1)

				return &testContext{
					fields: Fields{
						roomService:    roomService,
						messageService: messageService,
					},
					args: Args{
						ctx:   context.Background(),
						token: "a1b2c3",
					},
					returns: Returns{
						messages: messages,
						err:      nil,
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
				u := &messageUsecase{
					roomService:    tc.fields.roomService,
					messageService: tc.fields.messageService,
				}
				messages, err := u.GetAllMessages(tc.args.ctx, tc.args.token)
				if !reflect.DeepEqual(err, tc.returns.err) {
					t.Errorf("GetAllMessages() error = %v, wantErr %v", err, tc.returns.err)
				}
				if !reflect.DeepEqual(messages, tc.returns.messages) {
					t.Errorf("GetAllMessages() messages = %v, want %v", messages, tc.returns.messages)
				}
			})
		})
	}
}
