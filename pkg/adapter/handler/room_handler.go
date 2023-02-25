package handler

import (
	"github.com/isso-719/gaya-on-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type roomHandler struct {
	usecase usecase.IFRoomUsecase
}

func NewRoomHandler(su usecase.IFRoomUsecase) *roomHandler {
	return &roomHandler{
		usecase: su,
	}
}

type createRoomSuccessResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (sh *roomHandler) CreateRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		token, err := sh.usecase.CreateRoom(ctx)
		if err != nil {
			return internalServerErrorResponse(c, err)
		}
		return c.JSON(http.StatusOK, &createRoomSuccessResponse{
			Message: "success",
			Token:   token,
		})
	}
}

type findRoomRequest struct {
	Token string `json:"token"`
}
type findRoomResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (sh *roomHandler) FindRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req findRoomRequest
		if err := c.Bind(&req); err != nil {
			return internalServerErrorResponse(c, err)
		}
		token := req.Token
		ok, err := sh.usecase.FindRoom(ctx, token)
		if err != nil {
			return internalServerErrorResponse(c, err)
		}
		if !ok {
			return c.JSON(http.StatusBadRequest,
				&findRoomResponse{
					Message: "not found",
					Token:   token,
				})
		}
		return c.JSON(http.StatusOK, &findRoomResponse{
			Message: "found",
			Token:   token,
		})
	}
}

// TODO: JoinRoom : Start WebSocket connection
//func (sh *roomHandler) JoinRoom() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		ctx := c.Request().Context()
//		token := c.Param("token")
//		websocket.Handler(func(ws *websocket.Conn) {
//			defer ws.Close()
//			sh.usecase.JoinRoom(ctx, token, ws)
//		}).ServeHTTP(c.Response(), c.Request())
//		return nil
//	}
//}
