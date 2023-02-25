package handler

import (
	"github.com/isso-719/gaya-on-server/pkg/domain/model"
	"github.com/isso-719/gaya-on-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type messageHandler struct {
	usecase usecase.IFMessageUsecase
}

func NewMessageHandler(su usecase.IFMessageUsecase) *messageHandler {
	return &messageHandler{
		usecase: su,
	}
}

func roomNotFoundResponse(c echo.Context) error {
	return c.JSON(
		http.StatusBadRequest,
		&ErrorResponse{
			Message: "failed",
			Error:   "room not found",
		})
}

type createMessageRequest struct {
	Token       string `json:"token"`
	MessageType string `json:"message_type"`
	MessageBody string `json:"message_body"`
}
type createMessageResponse struct {
	Message string `json:"message"`
}

func (sh *messageHandler) CreateMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req createMessageRequest
		if err := c.Bind(&req); err != nil {
			return internalServerErrorResponse(c, err)
		}
		token := req.Token
		messageType := req.MessageType
		messageBody := req.MessageBody
		err := sh.usecase.CreateMessage(ctx, token, messageType, messageBody)
		if err != nil {
			if err.Error() == "room not found" {
				return roomNotFoundResponse(c)
			}
			return internalServerErrorResponse(c, err)
		}
		return c.JSON(
			http.StatusOK,
			&createMessageResponse{
				Message: "success",
			},
		)
	}
}

type getAllMessagesRequest struct {
	Token string `json:"token"`
}
type getAllMessagesResponse struct {
	Message  string `json:"message"`
	Messages []*model.MessageTypeAndBody
}

func (sh *messageHandler) GetAllMessages() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req getAllMessagesRequest
		if err := c.Bind(&req); err != nil {
			return internalServerErrorResponse(c, err)
		}
		token := req.Token
		messages, err := sh.usecase.GetAllMessages(ctx, token)
		if err != nil {
			if err.Error() == "room not found" {
				return roomNotFoundResponse(c)
			}
			return internalServerErrorResponse(c, err)
		}
		return c.JSON(
			http.StatusOK,
			&getAllMessagesResponse{
				Message:  "success",
				Messages: messages,
			},
		)
	}
}
