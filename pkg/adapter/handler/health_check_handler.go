package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type healthCheckResponse struct {
	HealthCheck string `json:"health_check"`
}

func HealthCheck(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		healthCheckResponse{HealthCheck: "ok"})
}
