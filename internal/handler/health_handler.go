package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
}

func (h *Handlers) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
