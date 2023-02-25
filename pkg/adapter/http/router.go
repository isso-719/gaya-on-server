package http

import (
	"github.com/isso-719/gaya-on-server/pkg/adapter/handler"
	"github.com/isso-719/gaya-on-server/pkg/domain/service"
	"github.com/isso-719/gaya-on-server/pkg/infra"
	"github.com/isso-719/gaya-on-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiVersion = "/v1"

	healthCheckRoot = "/health-check"
	roomRoot        = apiVersion + "/room"
	messageRoot     = apiVersion + "/message"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
		// TODO: CSRF対策
		// TODO: JWT認証
	)

	SQLConn := infra.NewSQLConnector()

	// Health Check Group
	healthCheckGroup := e.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, handler.HealthCheck)
	}

	// Room Group
	roomRepository := infra.NewRoomRepository(SQLConn.Conn)
	roomService := service.NewRoomService(roomRepository)
	roomUseCase := usecase.NewRoomUsecase(roomService)
	// Room Migration
	err := roomRepository.Migrate()
	if err != nil {
		panic(err)
	}

	roomGroup := e.Group(roomRoot)
	{
		handler := handler.NewRoomHandler(roomUseCase)

		var relativePath string

		relativePath = "/create"
		roomGroup.POST(relativePath, handler.CreateRoom())

		relativePath = "/find"
		roomGroup.POST(relativePath, handler.FindRoom())
	}

	// Message Group
	messageRepository := infra.NewMessageRepository(SQLConn.Conn)
	messageService := service.NewMessageService(messageRepository)
	messageUseCase := usecase.NewMessageUsecase(messageService, roomService)
	// Message Migration
	err = messageRepository.Migrate()
	if err != nil {
		panic(err)
	}

	messageGroup := e.Group(messageRoot)
	{
		handler := handler.NewMessageHandler(messageUseCase)

		relativePath := "/send"
		messageGroup.POST(relativePath, handler.CreateMessage())

		relativePath = "/get-all"
		messageGroup.POST(relativePath, handler.GetAllMessages())
	}

	return e
}
