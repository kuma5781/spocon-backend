package handler

import (
	"net/http"
	"spocon-backend/internal/openapi"

	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
}

func (h *Handlers) CreateTeam(c echo.Context) error {
	var body openapi.CreateTeamRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return nil
}
