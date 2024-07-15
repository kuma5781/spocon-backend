package handler

import (
	"net/http"
	"spocon-backend/internal/openapi"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h *Handlers) CreateUser(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) GetUser(c echo.Context, params openapi.GetUserParams) error {
	return c.NoContent(http.StatusOK)
}
