package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
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
		// TODO: 絵文字メッセージの追加テストの実装, 絵文字が一文字でないときは受け入れない時の実装
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
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			go func() {
				tc := tt.testContext(ctrl)
				s := &messageUsecase{
					roomService:    tc.fields.roomService,
					messageService: tc.fields.messageService,
				}
				err := s.CreateMessage(tc.args.ctx, tc.args.token, tc.args.messageType, tc.args.messageBody)
				if err != nil {
					t.Error(err)
				}

				cancel()
			}()
			<-ctx.Done()
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
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			go func() {
				tc := tt.testContext(ctrl)
				s := &messageUsecase{
					roomService:    tc.fields.roomService,
					messageService: tc.fields.messageService,
				}
				messages, err := s.GetAllMessages(tc.args.ctx, tc.args.token)
				if err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(messages, tc.returns.messages) {
					t.Errorf("messages = %v, want %v", messages, tc.returns.messages)
				}

				cancel()
			}()
			<-ctx.Done()
		})
	}
}
