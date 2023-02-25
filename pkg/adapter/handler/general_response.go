package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func internalServerErrorResponse(c echo.Context, err error) error {
	return c.JSON(
		http.StatusInternalServerError,
		&ErrorResponse{
			Message: "failed",
			Error:   err.Error(),
		})
}
