package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h *Handlers) CreateUser(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
